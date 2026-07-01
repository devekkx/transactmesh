package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	natspkg "github.com/devekkx/transactmesh/packages/nats"

	"github.com/devekkx/transactmesh/wallet-service/internal/events"
	httpapi "github.com/devekkx/transactmesh/wallet-service/internal/http"
	"github.com/devekkx/transactmesh/wallet-service/internal/wallet"
)

func main() {

	// NATS client
	bus, err := natspkg.New(natspkg.Config{
		URL:           "nats://localhost:4222",
		Name:          "wallet-service",
		MaxReconnects: -1,
		ReconnectWait: 2 * time.Second,
		Timeout:       5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	// repo (temporary in-memory)
	repo := wallet.NewInMemoryRepo()

	// event publisher
	publisher := events.NewPublisher(bus)

	// service
	svc := wallet.NewService(repo, publisher)

	// gin setup
	r := gin.New()

	// production middlewares
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})

	handler := httpapi.NewHandler(svc)
	handler.RegisterRoutes(r)

	log.Println("wallet-service running on :8080")
	log.Fatal(r.Run(":8080"))
}
