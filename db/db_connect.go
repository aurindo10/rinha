package db

import (
	"fmt"
	"rinha/config"
	"rinha/db/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectToDb() error {
	user := config.Config("POSTGRESQL_USERNAME")
	password := config.Config("POSTGRESQL_PASSWORD")
	host := config.Config("HOST")
	port := config.Config("PORT")
	db_name := config.Config("POSTGRESQL_DATABASE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",host,user,password,db_name,port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN: dsn,
	// 	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	//   }), &gorm.Config{})


	if err != nil {
		panic("failed to connect database")
	  }
	  if err = db.AutoMigrate(&schemas.Pessoas{}, &schemas.Stack{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	
	DB = db

	return nil
}