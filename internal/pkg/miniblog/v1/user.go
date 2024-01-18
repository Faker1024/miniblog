package v1

// CreateUserRequest 制定了`POST/v1/users`接口的请求参数
type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanumeric, required, string length(1|255)" `
	Password string `json:"password" valid:"required, string length(6|18)"`
	Nickname string `json:"nickname" valid:"required, string length(1|255)"`
	Email    string `json:"email" valid:"required, email"`
	Phone    string `json:"phone" valid:"required, string length(11|11)"`
}
