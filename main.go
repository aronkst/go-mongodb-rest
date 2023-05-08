package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aronkst/go-mongodb-rest/mongodb"
	"github.com/aronkst/go-mongodb-rest/web"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := getEnv("PORT")
	databaseName := getEnv("DATABASE")
	mongoConnection := getEnv("MONGO_CONNECTION")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnection))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(databaseName)
	mongoDB := mongodb.New(database)

	server := web.New(mongoDB)

	router := httprouter.New()

	router.POST("/query/:collection", server.Query)
	router.POST("/count/:collection", server.Count)
	router.POST("/paginate/:collection", server.Paginate)
	router.GET("/collection/:collection", server.List)
	router.GET("/collection/:collection/:_id", server.Show)
	router.POST("/collection/:collection", server.Insert)
	router.PUT("/collection/:collection", server.Replace)
	router.PATCH("/collection/:collection", server.Update)
	router.DELETE("/collection/:collection", server.Delete)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		err := fmt.Errorf("invalid %s environment variable", name)
		log.Fatal(err)
	}

	return value
}
