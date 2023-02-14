package response

type UserResponse struct {
	ID   string `json: "id"`
	Mail string `json: "mail"`
	Name string `json: "name"`
	Age  int8   `json: "age"`
}
