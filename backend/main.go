package main

import (
	"backend/config"
	"backend/domain/user"
	"backend/domain/workbench"
	"backend/infra"
	"context"

	"github.com/cv70/pkgo/mistake"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg, err := config.LoadConfig()
	mistake.Unwrap(err)

	// Initialize infrastructure with configuration
	registry, err := infra.NewRegistry(ctx, cfg)
	mistake.Unwrap(err)

	r := gin.Default()
	v1 := r.Group("/api/v1")

	// Register user routes
	userDomain := user.UserDomain{
		DB:    registry.DB,
		Redis: registry.Redis,
	}
	user.RegisterRoutes(v1, &userDomain)
	workbenchDomain, err := workbench.NewDomain(registry.DB)
	mistake.Unwrap(err)
	workbench.RegisterRoutes(v1, workbenchDomain)

	err = r.Run(":8888")
	mistake.Unwrap(err)
}
