package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/consts"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
	"github.com/xingrgx/WeShareX/utility/response"
)

type cChat struct {
	Users *gmap.Map
	Names *gset.StrSet
}

var Chat = cChat{
	Users: gmap.New(true),
	Names: gset.NewStrSet(true),
}

func (cc *cChat) IndexChatFriends(ctx context.Context, req *v1.IndexChatFriendsReq) (res *v1.IndexChatFriendsRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "通讯录",
	})
	return
}

func (cc *cChat) IndexChatRoom(ctx context.Context, req *v1.IndexChatRoomReq) (res *v1.IndexChatRoomRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "公共聊天室",
	})
	return
}

func (cc *cChat) Name(ctx context.Context, req *v1.ChatNameReq) (res *v1.ChatNameRes, err error) {
	var (
		session = g.RequestFromCtx(ctx).Session
	)
	// Create name in session.
	req.Name = service.Session().GetUser(ctx).Nickname
	g.Dump(req.Name)
	session.MustSet(consts.ChatSessionNameTemp, req.Name)
	if cc.Names.Contains(req.Name) {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Name)
	} else {
		session.MustSet(consts.ChatSessionName, req.Name)
		session.MustRemove(
			consts.ChatSessionNameTemp,
			consts.ChatSessionNameError,
		)
	}
	return
}

func (cc *cChat) Websocket(ctx context.Context, req *v1.ChatWebsocketReq) (res *v1.ChatWebsocketRes, err error) {
	var (
		r       = g.RequestFromCtx(ctx)
		ws      *ghttp.WebSocket
		msg     model.ChatMsg
		name    string
		msgByte []byte
	)
	if ws, err = r.WebSocket(); err != nil {
		g.Log().Error(ctx, err)
		return
	}

	name = service.Session().GetUser(ctx).Nickname

	// Create session data for current websocket.
	cc.Names.Add(name)
	cc.Users.Set(ws, name)

	// It notifies all clients that this websocket is online.
	if err = cc.writeGroupWithTypeList(); err != nil {
		return nil, err
	}

	for {
		// Blocking reading message from current websocket.
		_, msgByte, err = ws.ReadMessage()
		if err != nil {
			// Remove session info.
			cc.Names.Remove(name)
			cc.Users.Remove(ws)
			// It notifies all clients that this websocket is offline.
			_ = cc.writeGroupWithTypeList()
			return nil, nil
		}
		// Message decoding.
		if err = gjson.DecodeTo(msgByte, &msg); err != nil {
			_ = cc.write(ws, model.ChatMsg{
				Type: consts.ChatTypeError,
				Data: fmt.Sprintf(`invalid message: %s`, err.Error()),
				From: "",
			})
			continue
		}
		msg.From = name

		g.Log().Print(ctx, msg)

		switch msg.Type {
		case consts.ChatTypeSend:
			// Checks sending interval limit.
			var (
				cacheKey = fmt.Sprintf("ChatWebSocket:%p", ws)
			)
			if ok, _ := gcache.SetIfNotExist(ctx, cacheKey, struct{}{}, consts.ChatIntervalLimit); !ok {
				_ = cc.write(ws, model.ChatMsg{
					Type: consts.ChatTypeError,
					Data: `Message sending too frequently, why not a rest first`,
					From: "",
				})
				continue
			}
			// When new message retrieved, it notifies all clients.
			if msg.Data != nil {
				if err = cc.writeGroup(model.ChatMsg{
					Type: consts.ChatTypeSend,
					Data: ghtml.SpecialChars(gconv.String(msg.Data)),
					From: ghtml.SpecialChars(msg.From),
				}); err != nil {
					g.Log().Error(ctx, err)
				}
			}
		}
	}
}

// write sends message to current client.
func (cc *cChat) write(ws *ghttp.WebSocket, msg model.ChatMsg) error {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	return ws.WriteMessage(ghttp.WsMsgText, msgBytes)
}

// writeGroup sends message to all clients.
func (cc *cChat) writeGroup(msg model.ChatMsg) error {
	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	cc.Users.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			_ = user.(*ghttp.WebSocket).WriteMessage(ghttp.WsMsgText, []byte(b))
		}
	})

	return nil
}

// writeGroupWithTypeList sends "list" type message to all clients that can update users list in each client.
func (cc *cChat) writeGroupWithTypeList() error {
	array := garray.NewSortedStrArray()
	cc.Names.Iterator(func(v string) bool {
		array.Add(v)
		return true
	})
	if err := cc.writeGroup(model.ChatMsg{
		Type: consts.ChatTypeList,
		Data: array.Slice(),
		From: "",
	}); err != nil {
		return err
	}
	return nil
}

func (cc *cChat) ListFriends(ctx context.Context, req *v1.ListFriendsReq) (res *v1.ListFriendsRes, err error) {
	friends, err := service.Chat().GetAllFriends(service.Session().GetUser(ctx).Id)
	service.View().Render(ctx, model.View{
		Title: "通讯录",
		Data: g.Map {
			"Friends": friends,
		},
	})
	return
}

func (cc *cChat) AddFriend(ctx context.Context, req *v1.AddFriendReq) (res *v1.AddFriendRes, err error) {
	uid := service.Session().GetUser(ctx).Id
	err = service.Chat().AddFriend(ctx, uid, gconv.Uint(req.Id))
	if err != nil {
		return
	}
	return
}

func (cc *cChat) SearchFriend(ctx context.Context, req *v1.SearchFriendReq) (res *v1.SearchFriendRes, err error) {
	if friends, _ := service.User().GetUserByPassport(ctx, req.Msg); friends != nil {
		response.Json(g.RequestFromCtx(ctx), 0, "", friends)
		return
	}
	if friends, _ := service.User().GetUserByNickname(ctx, req.Msg); friends != nil {
		response.Json(g.RequestFromCtx(ctx), 0, "", friends)
		return
	}
	return res, gerror.New("查无此用户")
}

func (cc *cChat) Agree(ctx context.Context, req *v1.AgreeReq) (res *v1.AgreeRes, err error) {
	err = service.Chat().SetStatusTo2ById(ctx, service.Session().GetUser(ctx).Id, gconv.Uint(req.Id))
	return
}

func (cc *cChat) IndexChat(ctx context.Context, req *v1.ChatIndexReq) (res *v1.ChatIndexRes, err error) {
	f, err := service.User().GetUserProfileByID(ctx, req.Id)
	if err != nil {
		return
	} else {
		service.View().Render(ctx, model.View{
			Title: "好友通讯",
			Data: g.Map {
				"Friend": f,
			},
		})
		return
	}
}