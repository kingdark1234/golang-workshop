package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DatabaseServer ...
type DatabaseServer struct {
	gorm.Model
	Code  string
	Price uint
}

// NewDatabaseServer ...
func NewDatabaseServer() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=account password=postgres sslmode=disable")
	// db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=dvdrental password=12345")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&DatabaseServer{})

	// Create
	db.Create(&DatabaseServer{Code: "L1212", Price: 1000})

	// Read
	var product DatabaseServer
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
