package response

type PageList struct {
	PageSize int         `json:"pageSize"`
	PageNo   int         `json:"pageNo"`
	Total    int64       `json:"total"`
	DataList interface{} `json:"dataList"`
}
