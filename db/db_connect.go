package db

import (
	"fmt"
	"rinha/config"
	"rinha/db/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
func ConnectToDb() error {
	user := config.Config("POSTGRESQL_USERNAME")
	password := config.Config("POSTGRESQL_PASSWORD")
	host := config.Config("HOST")
	port := config.Config("PORT")
	db_name := config.Config("POSTGRESQL_DATABASE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",host,user,password,db_name,port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), 
	})

	if err != nil {
		panic("failed to connect database")
	  }
	  if err = db.AutoMigrate(&schemas.Pessoas{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	DB = db
	// db.Migrator().DropTable(&schemas.Pessoas{})

	return nil
}