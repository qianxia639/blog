package response

type Search struct {
	Id          interface{} `json:"id"`
	UserId      interface{} `json:"userId"`
	TypeId      interface{} `json:"typeId"`
	TypeName    string      `json:"typeName"`
	Username    string      `json:"username"`
	Title       interface{} `json:"title"`
	Description interface{} `json:"description"`
	UpdatedAt   string      `json:"updatedAt"`
	Tags        interface{} `json:"tags"`
}
