package main

import (
	"log"
	"purchase-price-calculation/internal/config"
	"purchase-price-calculation/internal/db"
	"purchase-price-calculation/internal/http/router"
)

func main() {
	cfg := config.Load()
	if err := db.Init(cfg); err != nil {
		log.Fatal(err)
	}

	r := router.New(cfg)
	log.Printf("Gin server listening on %s", cfg.Addr)
	if err := r.Run(cfg.Addr); err != nil {
		log.Fatal(err)
	}
}
