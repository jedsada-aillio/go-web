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
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jedsa-aillio/go-web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(os.Getenv(
		"MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx,
		collection)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	// router.PUT("/recipes/:id", UpdateRecipeHandler)
	// router.DELETE("/recipes/:id", DeleteRecipeHandler)
	// router.GET("/recipes/search", SearchRecipesHandler)
	router.Run()
}
