package service

import (
	"context"

	"choice-tech-project/internal/consts"
	"choice-tech-project/internal/model"
	"choice-tech-project/internal/repository"
)

const redisKey = consts.RedisKey
const redisTTL = consts.RedisCacheTTL // Cache expiration duration

// Service provides business logic for importing, viewing, updating, and deleting records.
type Service struct {
	MySQLRepo *repository.MySQLRepository // MySQL repository for DB operations
	RedisRepo *repository.RedisRepository // Redis repository for caching
}

// NewService creates a new Service instance.
func NewService(mysqlRepo *repository.MySQLRepository, redisRepo *repository.RedisRepository) *Service {
	return &Service{
		MySQLRepo: mysqlRepo,
		RedisRepo: redisRepo,
	}
}

// ImportRecords inserts records into MySQL after clearing the table and resetting the auto-increment ID.
func (s *Service) SaveRecords(ctx context.Context, records []model.Record) error {
	// Truncate the table to remove all data and reset auto-increment
	err := s.MySQLRepo.TruncateTable(ctx)
	if err != nil {
		return err
	}
	// Insert new records
	err = s.MySQLRepo.InsertRecordsWithGoroutines(ctx, records)
	if err != nil {
		return err
	}
	return s.RedisRepo.SetRecords(ctx, consts.RedisKey, records, consts.RedisCacheTTL)
}

// GetRecords returns all records, using Redis cache if available.
func (s *Service) GetRecords(ctx context.Context) ([]model.Record, error) {
	return s.RedisRepo.GetOrFetchRecords(ctx, redisKey, func() ([]model.Record, error) {
		return s.MySQLRepo.GetAllRecords(ctx)
	})
}

// UpdateRecord updates a record in MySQL and refreshes the Redis cache.
func (s *Service) UpdateRecord(ctx context.Context, rec model.Record) error {
	if err := s.MySQLRepo.UpdateRecord(ctx, rec); err != nil {
		return err
	}
	records, err := s.MySQLRepo.GetAllRecords(ctx)
	if err == nil {
		_ = s.RedisRepo.SetRecords(ctx, redisKey, records, redisTTL)
	}
	return nil
}

// DeleteRecord deletes a record in MySQL and refreshes the Redis cache.
func (s *Service) DeleteRecord(ctx context.Context, id int) error {
	if err := s.MySQLRepo.DeleteRecord(ctx, id); err != nil {
		return err
	}
	records, err := s.MySQLRepo.GetAllRecords(ctx)
	if err == nil {
		_ = s.RedisRepo.SetRecords(ctx, redisKey, records, redisTTL)
	}
	return nil
}

// GetRecordByID returns a single record by its ID, using Redis cache if available.
func (s *Service) GetRecordByID(ctx context.Context, id int) (*model.Record, error) {
	// Try to get all records from cache first
	records, err := s.RedisRepo.GetOrFetchRecords(ctx, consts.RedisKey, func() ([]model.Record, error) {
		return s.MySQLRepo.GetAllRecords(ctx)
	})
	if err == nil {
		for _, rec := range records {
			if rec.ID == id {
				return &rec, nil
			}
		}
	}
	// Fallback to DB direct fetch if not found in cache
	return s.MySQLRepo.GetRecordByID(ctx, id)
}
