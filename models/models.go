package models

import "time"

type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	*Request
	RateLimit      int           `json:"rate_limit"`
	ResetRateLimit time.Duration `json:"rate_limit_reset"`
}
