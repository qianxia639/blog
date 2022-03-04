package response

type PageList struct {
	Total       int64       `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	DataList    interface{} `json:"dataList"`
}
