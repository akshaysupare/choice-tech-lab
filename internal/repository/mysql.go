package repository

import (
	"context"
	"fmt"
	"log"
	"sync"

	"choice-tech-project/internal/model"
	"choice-tech-project/internal/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLRepository provides methods for interacting with the MySQL database using GORM.
type MySQLRepository struct {
	DB *gorm.DB
}

// NewMySQLRepository creates a new MySQLRepository with the given DSN using GORM.
func NewMySQLRepository(dsn string) (*MySQLRepository, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MySQLRepository{DB: db}, nil
}

// CreateTable auto-migrates the Record model to create the table if it does not exist.
func (r *MySQLRepository) CreateTable() error {
	return r.DB.AutoMigrate(&model.Record{})
}

// InsertRecords inserts records in batches of 100 for efficiency.
func (r *MySQLRepository) InsertRecords(ctx context.Context, records []model.Record) error {
	batchSize, concurrency := utils.DecideBatchSizeAndConcurrency(len(records))
	log.Printf("Inserting %d records in batches of %d with concurrency %d", len(records), batchSize, concurrency)
	return r.DB.WithContext(ctx).CreateInBatches(records, batchSize).Error
}
// InsertRecords inserts records in batches using multiple goroutines for concurrency.
func (r *MySQLRepository) InsertRecordsWithGoroutines(ctx context.Context, records []model.Record) error {
    batchSize, concurrency := utils.DecideBatchSizeAndConcurrency(len(records))
    log.Printf("Inserting %d records in batches of %d with concurrency %d", len(records), batchSize, concurrency)

    total := len(records)
    if total == 0 {
        log.Println("No records to insert.")
        return nil
    }

    // Channel to send batches to workers
    batchChan := make(chan []model.Record)
    // Channel to collect errors from workers
    errChan := make(chan error, concurrency)
    var wg sync.WaitGroup

    // Start worker goroutines
    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for batch := range batchChan {
                if err := r.DB.WithContext(ctx).CreateInBatches(batch, len(batch)).Error; err != nil {
                    errChan <- err
                }
            }
        }()
    }

    // Send batches to workers
    for i := 0; i < total; i += batchSize {
        end := i + batchSize
        if end > total {
            end = total
        }
        batchChan <- records[i:end]
    }
    close(batchChan)

    // Wait for all workers to finish
    wg.Wait()
    close(errChan)

    // Collect errors if any
    if len(errChan) > 0 {
        return fmt.Errorf("insertion encountered %d errors (check logs)", len(errChan))
    }

    return nil
}

// GetAllRecords fetches all records from the database.
func (r *MySQLRepository) GetAllRecords(ctx context.Context) ([]model.Record, error) {
	var records []model.Record
	result := r.DB.WithContext(ctx).Find(&records)
	return records, result.Error
}

// UpdateRecord updates a record by ID.
func (r *MySQLRepository) UpdateRecord(ctx context.Context, rec model.Record) error {
	return r.DB.WithContext(ctx).Model(&model.Record{}).Where("id = ?", rec.ID).Updates(rec).Error
}

// DeleteRecord deletes a record by ID.
func (r *MySQLRepository) DeleteRecord(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&model.Record{}, id).Error
}

// GetRecordByID fetches a single record by its ID from the database.
func (r *MySQLRepository) GetRecordByID(ctx context.Context, id int) (*model.Record, error) {
	var rec model.Record
	result := r.DB.WithContext(ctx).First(&rec, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rec, nil
}

// TruncateTable removes all data from the records table and resets the auto-increment ID.
func (r *MySQLRepository) TruncateTable(ctx context.Context) error {
	return r.DB.WithContext(ctx).Exec("TRUNCATE TABLE records").Error
}
