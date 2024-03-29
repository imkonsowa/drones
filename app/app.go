package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"drones/app/data/models"
	"drones/pkg/config"
)

type App struct {
	Config *config.Config
	Engine *gin.Engine
	DB     *gorm.DB
}

func NewApp() *App {
	config.Construct()
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=GMT",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
		cfg.DB.SSlMode,
	)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(fmt.Sprintf("failed to connect to DB; %s", dbErr))
	}
	migrateErr := db.AutoMigrate(&models.Drone{}, &models.Medication{})
	if migrateErr != nil {
		panic(fmt.Sprintf("failed to migrate; %s", migrateErr))
	}

	engine := gin.Default()

	_ = engine.SetTrustedProxies([]string{})

	return &App{
		Engine: engine,
		Config: cfg,
		DB:     db,
	}
}

func (a *App) Run() {
	if a.Engine == nil {
		panic("server engine is not constructed yet")
	}

	cfg := config.GetConfig()

	err := a.Engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
	if err != nil {
		return
	}
}
