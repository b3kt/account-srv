package vo

import "time"

// ResultMsg - Base result message format
type ResultMsg struct {
	Message    string `json:"message"`
	Error      bool   `json:"error"`
	StatusCode int    `json:"status_code"`
	Timestamp  time.Time
}
