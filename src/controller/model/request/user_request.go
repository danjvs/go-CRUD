package request

type UserRequest struct {
	Mail     string `json: "mail"`
	password string `json: "password"`
	Name     string `json: "name"`
	Age      int8   `json: "age"`
}
