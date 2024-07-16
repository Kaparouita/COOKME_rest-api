package repositories

import (
	"bufio"
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
		&models.Recipe{},
		&models.Order{},
		&models.Review{},
		&models.Keyword{},
		&models.MarketIngredient{},
	)

	// Insert the keywords
	// dbRepo := &Db{db}
	// err = dbRepo.insertKeywords("uniqueKeywords.txt")

	return db, nil
}

func (db *Db) insertKeywords(path string) error {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the file
	keywords := make([]models.Keyword, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyword := models.Keyword{
			Keyword: scanner.Text(),
		}
		keywords = append(keywords, keyword)
	}

	// Insert the keywords
	for _, keyword := range keywords {
		db.Create(&keyword)
	}

	return nil
}

func (db *Db) GetAllKeywords() []models.Keyword {
	var keywords []models.Keyword
	db.Find(&keywords)
	return keywords
}
