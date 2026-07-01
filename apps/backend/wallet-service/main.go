package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/devekkx/transactmesh/wallet-service/internal/config"
	"github.com/devekkx/transactmesh/wallet-service/internal/observability"
	"github.com/devekkx/transactmesh/wallet-service/internal/wallet"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	shutdown, err := observability.InitTracer(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = shutdown(ctx)
	}()

	walletService := wallet.NewService()

	mux := http.NewServeMux()

	mux.HandleFunc("/debit", func(w http.ResponseWriter, r *http.Request) {
		_ = walletService.Debit(r.Context(), "123", 100)
		w.Write([]byte("ok"))
	})

	handler := observability.NewHTTPHandler(mux, "wallet.http")

	log.Println("wallet-service running on :8080")
	_ = http.ListenAndServe(":8080", handler)
}
