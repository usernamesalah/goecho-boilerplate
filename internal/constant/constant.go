package constant

import "time"

// List of internal constant
const (
	DBStringConnection = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4"

	WriteTimeout = 10 * time.Second
	JsonHeader   = "application/json"
)
