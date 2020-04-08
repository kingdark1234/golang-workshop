package main

import "workshop-service/internals/container"

func main() {
	// func main() {
	// 	configs.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=robodev dbname=go-workshop password=1qazxsw2 sslmode=disable")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer configs.DB.Close()

	// 	// migrate db
	// 	configs.DB.AutoMigrate(&models.Author{})

	// 	r := routers.SetupRouter()
	// 	// Listen and Server in 0.0.0.0:8080
	// 	r.Run(":8080")
	// }
	c, err := container.NewContainer()
	if err != nil {
		panic(err)
	}
	if err := c.Start(); err != nil {
		panic(err)
	}
}
