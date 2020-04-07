package database

import (
	"fmt"
	"log"
	"workshop-service/internals/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ServerBase ...
type ServerBase struct {
	DB *gorm.DB
}

// Initialize ...
func (server *ServerBase) Initialize() {
	var db = server.DB
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=people password=postgres sslmode=disable")
	if err != nil {
		fmt.Println("Cannot connect to postgrest database")
		log.Fatal("This is the error:", err)
		defer db.Close()
		panic("Failed to connected database.")
	} else {
		fmt.Println("We are connected to the postgrest database")
		db.Debug().AutoMigrate(&entity.People{}) //database migration
	}
}

// NewDatabaseServer ...
func NewDatabaseServer() *ServerBase {
	return &ServerBase{}
}
