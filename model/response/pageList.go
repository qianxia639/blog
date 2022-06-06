package response

type PageList struct {
	PageSize int         `json:"pageSize,omitempty"`
	PageNo   int         `json:"pageNo,omitempty"`
	Total    int64       `json:"total,omitempty"`
	DataList interface{} `json:"dataList,omitempty"`
}
