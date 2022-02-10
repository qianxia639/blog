package vo

type PageListVO struct {
	Total    int64       `json:"total"`
	DataList interface{} `json:"dataList"`
}
