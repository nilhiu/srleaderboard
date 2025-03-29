package ajax

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddRunRequest struct {
	Time Duration `json:"time"`
}

type ValidateTimeRequest struct {
	Time string `json:"time"`
}
