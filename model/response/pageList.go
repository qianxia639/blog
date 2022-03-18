package response

type PageList struct {
	Total    int64 `json:"total"`
	PageSize int   `json:"pageSize"`
	PageNum  int   `json:"pageNum"`

	DataList interface{} `json:"dataList"`
}
