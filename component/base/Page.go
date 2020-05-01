package base

type Page struct {
	First     bool        `json:"first,omitempty"`
	Last      bool        `json:"last,omitempty"`
	PageNo    int         `json:"pageNo,omitempty"`
	PageSize  int         `json:"pageSize,omitempty"`
	Total     int64       `json:"total,omitempty"`
	TotalPage int         `json:"totalPage,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func (p *Page) Offset() int64 {
	return int64((p.PageNo - 1) * p.PageSize)
}
