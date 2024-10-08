package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-cleanarch/pkg/factory/config"

	"go.uber.org/zap"
)

func ConnTotDB(logger *zap.Logger) *gorm.DB {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ", cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	logger.Info("Database connected successfully")

	// Migrate the schema LocationTable
	if err := db.AutoMigrate(&LocationTable{}, &TempleLocList{}, &SubLocList{}, &VisitLog{}, &ArtLocList{}, &ArtEvent{}, &ArtSubEvent{}, &TbMap{}); err != nil {
		logger.Error("Error migrating schema", zap.Error(err))
		panic(err)
	}
	logger.Info("Database schema migrated successfully")

	return db
}
