package migrations

import (
	"go_todo/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MigrateDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/go_todos?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// AutoMigrate the schema
	err = db.AutoMigrate(&models.TodoList{}, &models.Todo{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db
}
