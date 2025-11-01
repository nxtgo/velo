package main

import (
	"log"

	"velo/internal/config"
	"velo/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	log.Printf("velo listening on %s\n", cfg.Addr)
	srv.ListenAndServe()
}
