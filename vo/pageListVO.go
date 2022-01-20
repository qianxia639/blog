package vo

type PageListVO struct {
	Total    int         `json:"total"`
	DataList interface{} `json:"dataList"`
}
