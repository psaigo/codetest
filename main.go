package main

import (
	"crypto-service/handlers"
	"crypto-service/internal/client"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const url = "https://api.hitbtc.com/api/2/public"

func main() {

	client := client.NewClient(url)
	router := handlers.SetupRouter(handlers.Service{Client: client})
	port := "0.0.0.0:8080"
	listenAndServe(router, port)
}

func listenAndServe(router *gin.Engine, port string) {

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
