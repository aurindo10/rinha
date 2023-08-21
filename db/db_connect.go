package db

import (
	"fmt"
	"os"
	"rinha/db/schemas"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDb() error {
	user := os.Getenv("POSTGRESQL_USERNAME")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	db_name := os.Getenv("POSTGRESQL_DATABASE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, db_name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// Configurando o pool de conexão
	sqlDB, dbErr := db.DB() // Obtem o sql.DB subjacente
	if dbErr != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %v", dbErr)
	}

	// Configurando as conexões
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)

	if err = db.AutoMigrate(&schemas.Pessoas{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	DB = db
	// db.Migrator().DropTable(&schemas.Pessoas{})

	return nil
}
