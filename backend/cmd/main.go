package main

import (
	"context"
	"log"
	"time"

	"github.com/guilhermefill/FinTrak/cmd/web"
	"github.com/guilhermefill/FinTrak/pkg/models"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := models.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer models.DbClient.Close()

	srv := web.Server()
	srv.ListenAndServe()
}
