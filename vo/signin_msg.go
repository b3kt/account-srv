package vo

import (
	"time"
)

// SigninRequestMsg - request format for signup
type SigninRequestMsg struct {
	// Email           string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SigninResponseMsg - response format
type SigninResponseMsg struct {
	Message     string    `json:"message"`
	ErrorCode   string    `json:"error_code"`
	StatusCode  int       `json:"status_code"`
	Timestamp   time.Time `json:"timestamp"`
	AccessToken string    `json:"token"`
	Username    string    `json:"username"`
}
