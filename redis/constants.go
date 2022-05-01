package redis

import "time"

const (
	maxRetries      = 5
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 500 * time.Millisecond
	dialTimeout     = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 3 * time.Second
	minIdleConns    = 20
	poolTimeout     = 6 * time.Second
	idleTimeout     = 12 * time.Second
)
