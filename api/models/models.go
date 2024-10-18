package models

import "time"

type ShortUrlRequest struct {
	URL             string        `json:"url"`
	CustomShortCode string        `json:"custom_url"`
	Expiry          time.Duration `json:"expiry"`
}

type ShortUrlResponse struct {
	URL             string        `json:"url"`
	CustomShortCode string        `json:"custom_url"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}
