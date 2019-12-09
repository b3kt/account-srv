package vo

// SignupRequestMsg - request format for signup
type SignupRequestMsg struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Confirmation    bool   `json:"confirmation"`
}

// SignupResponseMsg - response format
type SignupResponseMsg struct {
	Header ResultMsg `json:"header"`
}
