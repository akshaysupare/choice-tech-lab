package repository

import (
	"context"

	"choice-tech-project/internal/consts"
	"choice-tech-project/internal/model"

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
	return r.DB.WithContext(ctx).CreateInBatches(records, consts.BatchSize).Error
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
