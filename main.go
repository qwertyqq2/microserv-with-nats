package main

import (
	"L0task/broker"
	"L0task/controller"
	"L0task/store"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	store, err := store.NewStore(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()

	fmt.Println(os.Getenv("CLASTER_ID"))
	fmt.Println(os.Getenv("NATS_URL"))
	fmt.Println("........................")
	if err := controller.Registration(router, store, broker.Config{
		ClusterID: os.Getenv("CLASTER_ID"),
		URL:       os.Getenv("NATS_URL"),
		Subject:   os.Getenv("SUBJECT"),
	}); err != nil {
		log.Fatal(err)
	}

	router.Run(":8000")

}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
