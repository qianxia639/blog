package response

type PageList struct {
	Pagination struct {
		Total       int64 `json:"total"`
		PerPage     int   `json:"per_page"`
		CurrentPage int   `json:"current_page"`
		LastPage    int   `json:"last_page"`
	} `json:"pagination"`

	DataList interface{} `json:"dataList"`
}
