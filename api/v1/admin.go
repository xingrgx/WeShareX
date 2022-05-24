package v1

import "github.com/gogf/gf/v2/frame/g"

type IndexAdminReq struct {
	g.Meta `path:"/admin" method:"get"`
}

type IndexAdminRes struct {

}

type AdminIndexUsersReq struct {
	g.Meta `path:"/admin/users" method:"get"`
}

type AdminIndexUsersRes struct {

}

type EnableReq struct { 
	g.Meta `path:"/admin/enable" method:"post"`
	UserId uint `json:"userId"`
}

type EnableRes struct {

}

type DisableReq struct { 
	g.Meta `path:"/admin/disable" method:"post"`
	UserId uint `json:"userId"`
}

type DisableRes struct {

}