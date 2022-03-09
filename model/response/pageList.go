package response

type PageList struct {
	Total       int64       `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	DataList    interface{} `json:"dataList"`
}
