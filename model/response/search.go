package response

type Search struct {
	Id        interface{} `json:"id"`
	UserId    interface{} `json:"userId"`
	TypeId    interface{} `json:"typeId"`
	TypeName  string      `json:"typeName"`
	Nickname  string      `json:"nickname"`
	Title     interface{} `json:"title"`
	Content   interface{} `json:"content"`
	UpdatedAt interface{} `json:"updatedAt"`
	Tags      interface{} `json:"tags"`
}
