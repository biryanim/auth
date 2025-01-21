package main

import (
	"context"
	"github.com/biryanim/auth/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
