package consts

import "time"

const (
	BatchSize     = 100
	RedisKey      = "imported_records"
	RedisCacheTTL = 5 * time.Minute
)
