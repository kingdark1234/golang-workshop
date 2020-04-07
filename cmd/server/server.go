package main

import (
	// databaseServer "golang-workshop/src/infrastructure/database"
	httpServer "workshop-service/internals/infrastructure/http"
)

func main() {
	sv := httpServer.NewServerHTTP()
	// databaseServer.NewDatabaseServer()
	sv.Start()
}
