package vo

import (
	"time"

	"github.com/Nerzal/gocloak/v3"
)

// RecoveryRequestMsg - request format for recovery
type RecoveryRequestMsg struct {
	Email string `json:"email"`
}

// RecoveryResponseMsg - response format
type RecoveryResponseMsg struct {
	AccessToken *gocloak.JWT `json:"jwt"`
	Header      ResultMsg    `json:"header"`
	Body        struct {
		RecoveryToken string    `json:"recovery_token"`
		Expire        time.Time `json:"expire"`
	} `json:"body"`
}
