// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//	Schemes: http, https
//
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: jedsada@aillio.com
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"go_backend/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

var recipesHandler *handlers.RecipesHandler

func init() {
	client := recipesHandler.Influxdb_connect()
	// Check for an error
	if client != nil {
		log.Println("Connected to InfluxDB")
	}
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipeHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.GET("/recipes/search", recipesHandler.SearchRecipeHandler)
	router.Run()

	recipesHandler.Influxdb_connect().Close()
}
