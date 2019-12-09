package vo

import (
	"github.com/Nerzal/gocloak/v3"
)

// SigninRequestMsg - request format for signup
type SigninRequestMsg struct {
	// Email           string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SigninResponseMsg - response format
type SigninResponseMsg struct {
	AccessToken *gocloak.JWT `json:"jwt"`
	Header      ResultMsg    `json:"header"`
	Body        struct {
		Username string `json:"username"`
	} `json:"body"`
}
