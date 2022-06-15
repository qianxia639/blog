package response

type Search struct {
	Id        interface{} `json:"id"`
	UserId    interface{} `json:"userId"`
	TypeId    interface{} `json:"typeId"`
	TypeName  string      `json:"typeName"`
	Username  string      `json:"username"`
	Title     interface{} `json:"title"`
	Content   interface{} `json:"content"`
	UpdatedAt interface{} `json:"updatedAt"`
	Tags      interface{} `json:"tags"`
}
