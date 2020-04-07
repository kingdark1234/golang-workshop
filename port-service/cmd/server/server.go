package main

import (
	databaseServer "port-service/internals/infrastructure/database"
	httpServer "port-service/internals/infrastructure/http"
)

func main() {
	sv := httpServer.NewServerHTTP()
	databaseServer.NewDatabaseServer()
	sv.Start()
}
