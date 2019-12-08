package vo

import "time"

// SignupRequestMsg - request format for signup
type SignupRequestMsg struct {
	// Header struct {
	// 	ClientID     string `json:"client_id"`
	// 	ClientSecret string `json:"client_secret"`
	// } `json:"header"`
	// Content struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Confirmation    bool   `json:"confirmation"`
	// } `json:"content"`
}

// SignupResponseMsg - response format
type SignupResponseMsg struct {
	// Header struct {
	Message    string    `json:"message"`
	ErrorCode  string    `json:"error_code"`
	StatusCode int       `json:"status_code"`
	Timestamp  time.Time `json:"timestamp"`
	// } `json:"header"`
	// Content struct {
	// 	FirstName       string `json:"first_name"`
	// 	LastName        string `json:"last_name"`
	// 	Username        string `json:"username"`
	// 	Password        string `json:"password"`
	// 	PasswordConfirm string `json:"password_confirm"`
	// 	Confirmation    bool   `json:"confirmation"`
	// } `json:"content"`
}
