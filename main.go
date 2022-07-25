package main

import (
	"github.com/letiercamila/api-go-rest/database"
	"github.com/letiercamila/api-go-rest/routes"
)

func main() {
	database.ConectDB()
	routes.HandleRequest()
}
