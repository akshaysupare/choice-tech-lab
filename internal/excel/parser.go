package excel

import (
	"errors"
	"fmt"

	"choice-tech-project/internal/model"

	"github.com/xuri/excelize/v2"
)

// expectedHeaders defines the required column headers for the Excel import.
var expectedHeaders = []string{
	"first_name", "last_name", "company_name", "address", "city",
	"county", "postal", "phone", "email", "web",
}

// ParseExcelFile parses the Excel file at filePath, validates the header row, and returns the records.
// Only the header row is validated; data rows are imported as-is.
func ParseExcelFile(filePath string) ([]model.Record, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, errors.New("no data rows found")
	}
	// Header validation only
	for i, h := range expectedHeaders {
		if i >= len(rows[0]) || rows[0][i] != h {
			return nil, fmt.Errorf("invalid header at column %d: expected '%s', got '%s'", i+1, h, rows[0][i])
		}
	}
	var records []model.Record
	for _, row := range rows[1:] {
		if len(row) < len(expectedHeaders) {
			continue // skip incomplete rows
		}
		records = append(records, model.Record{
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			County:      row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		})
	}
	return records, nil
}
