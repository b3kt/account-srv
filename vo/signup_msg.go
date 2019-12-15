package vo

// SignupRequestMsg - request format for signup
type SignupRequestMsg struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Confirmation    bool   `json:"isConfirmed"`
}

// SignupResponseMsg - response format
type SignupResponseMsg struct {
	Header ResultMsg `json:"header"`
}
