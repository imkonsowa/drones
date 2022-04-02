package app

import (
	"drones/app/http/middlewares"
	"drones/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

type App struct {
	Config *config.Config
	Engine *gin.Engine
}

func NewApp() *App {
	config.Construct()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env vars. Err: %s", err)
	}

	engine := gin.Default()
	engine.Use(middlewares.JsonResponseHeader)

	engine.SetTrustedProxies([]string{})

	return &App{
		Engine: engine,
		Config: config.GetConfig(),
	}
}

func (a *App) Run() {
	if a.Engine == nil {
		panic("server engine is not constructed yet")
	}

	cfg := config.GetConfig()

	err := a.Engine.Run(fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		return
	}
}
