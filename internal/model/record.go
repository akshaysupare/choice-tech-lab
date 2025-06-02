package model

// Record represents a row from the Excel file with all columns
// from the provided sample. Each field maps to a column in the Excel sheet.
type Record struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName   string `json:"first_name" gorm:"column:first_name;type:varchar(100)"`
	LastName    string `json:"last_name" gorm:"column:last_name;type:varchar(100)"`
	CompanyName string `json:"company_name" gorm:"column:company_name;type:varchar(255)"`
	Address     string `json:"address" gorm:"column:address;type:varchar(255)"`
	City        string `json:"city" gorm:"column:city;type:varchar(100)"`
	County      string `json:"county" gorm:"column:county;type:varchar(100)"`
	Postal      string `json:"postal" gorm:"column:postal;type:varchar(50)"`
	Phone       string `json:"phone" gorm:"column:phone;type:varchar(50)"`
	Email       string `json:"email" gorm:"column:email;type:varchar(100)"`
	Web         string `json:"web" gorm:"column:web;type:varchar(255)"`
}
