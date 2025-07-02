package repo

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Transactions struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        time.Time `json:"date" gorm:"index"`
	Description string    `json:"description" gorm:"type:varchar(50)"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
}

var DB *gorm.DB

func SetupDatabase() {

	db, err := gorm.Open(sqlite.Open("transaction.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!!!")
	}

	err = db.AutoMigrate(&Transactions{})
	if err != nil {
		panic("Failed to migrate to database!!!")
	}

	DB = db
}
