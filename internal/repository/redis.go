package repository

import (
	"context"
	"encoding/json"
	"time"

	"choice-tech-project/internal/consts"
	"choice-tech-project/internal/model"

	"github.com/go-redis/redis/v8"
)

// RedisRepository provides methods for interacting with Redis for caching records.
type RedisRepository struct {
	Client *redis.Client
}

// NewRedisRepository creates a new RedisRepository and checks the connection.
func NewRedisRepository(addr, password string, db int) (*RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisRepository{Client: client}, nil
}

// SetRecords stores records in Redis with the given key and expiration.
func (r *RedisRepository) SetRecords(ctx context.Context, key string, records []model.Record, expiration time.Duration) error {
	data, err := json.Marshal(records)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, data, expiration).Err()
}

// GetRecords retrieves records from Redis by key.
func (r *RedisRepository) GetRecords(ctx context.Context, key string) ([]model.Record, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var records []model.Record
	if err := json.Unmarshal([]byte(val), &records); err != nil {
		return nil, err
	}
	return records, nil
}

// DeleteRecords removes records from Redis by key.
func (r *RedisRepository) DeleteRecords(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// GetOrFetchRecords tries to get records from Redis, and if not found, fetches from DB and caches them.
func (r *RedisRepository) GetOrFetchRecords(ctx context.Context, key string, fetchFromDB func() ([]model.Record, error)) ([]model.Record, error) {
	expiration := consts.RedisCacheTTL
	records, err := r.GetRecords(ctx, key)
	if err == nil && len(records) > 0 {
		return records, nil
	}
	// If not found in Redis, fetch from DB
	records, err = fetchFromDB()
	if err != nil {
		return nil, err
	}
	_ = r.SetRecords(ctx, key, records, expiration)
	return records, nil
}
