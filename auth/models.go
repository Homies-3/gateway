package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Company  string `json:"company"`
}
