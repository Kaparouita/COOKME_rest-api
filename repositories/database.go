package repositories

import (
	"fmt"
	"log"
	"os"

	"rest-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	*gorm.DB
}

func NewDbRepo() *Db {
	db, err := connectDb()
	if err != nil {
		log.Fatal(err)
	}
	return &Db{
		db,
	}
}

func connectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Mingrations")

	//Create the tables
	db.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Recipe{},
	)

	return db, nil
}
