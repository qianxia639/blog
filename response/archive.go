package response

type Archive struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Flag  string `json:"flag"`
	Year  string `json:"year"`
	Date  string `json:"date"`
}
