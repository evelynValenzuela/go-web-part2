package main

import (
	"os"
	"parte2/cmd/server/routes"
	"parte2/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	//Instances
	store := store.NewStore(os.Getenv("FILEPATH"))

	products, err := store.ReadFile()

	if err != nil {
		panic(err)
	}

	server := gin.Default()
	router := routes.NewRouter(store, server, &products)
	router.SetRoutes()

	if err := server.Run(":8081"); err != nil {
		panic(err)
	}
}
