package response

type PageList struct {
	PageSize int   `json:"pageSize"`
	PageNum  int   `json:"pageNum"`
	Total    int64 `json:"total"`

	DataList interface{} `json:"dataList"`
}
