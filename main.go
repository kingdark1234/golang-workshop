package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"workshop-service/configs"
	"workshop-service/models"
	"workshop-service/routers"
)

var err error

func main() {
	configs.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=robodev dbname=go-workshop password=1qazxsw2 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer configs.DB.Close()

	// migrate db
	configs.DB.AutoMigrate(&models.Author{})

	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
