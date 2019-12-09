package vo

// ResetPassRequestMsg - request format for reset password
type ResetPassRequestMsg struct {
	Email                string `json:"email"`
	RecoveryToken        string `json:"recovery_token"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// ResetPassResponseMsg - response format
type ResetPassResponseMsg struct {
	Header ResultMsg `json:"header"`
}
