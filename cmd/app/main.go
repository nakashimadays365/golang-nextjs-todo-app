package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"todo/config"
	"todo/logger"
	"todo/repo"
	"todo/server"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	rp, err := repo.NewRepo(ctx, cfg)
	if err != nil {
		logger.Error(err)
		return
	}
	defer rp.DB.Close()

	srv := server.NewServer(rp)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Infof("server start :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv); err != nil {
		logger.Error(err)
		return
	}
}
