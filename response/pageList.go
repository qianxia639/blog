package response

type PageList struct {
	Total    int64       `json:"total"`
	DataList interface{} `json:"dataList"`
}
