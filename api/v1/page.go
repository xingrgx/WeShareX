package v1

type PaginationReq struct {
	Page int `json:"page" d:"1" dc:"分页码，默认为1"`
	Size int `json:"size" d:"10" dc:"分页数量，默认为10"`
}

type PaginationRes struct {
	Total int `dc:"总数"`
}