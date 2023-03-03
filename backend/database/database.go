package database

import (
	"errors"
	"fmt"
	"github.com/julyusmanurung/Kredit/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	db, err := callDB()

	if err != nil {
		return nil, err
	}

	db, err = checkConn(db)
	if err != nil {
		return nil, err
	}

	db, err = migrateDB(db)
	if err != nil {
		return nil, errorConn(err)
	}

	seedInitialIdTab(db)

	return db, nil
}

func errorConn(err error) error {
	return fmt.Errorf("connect database fails: %w", err)
}

func callDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	if err != nil {
		return nil, errorConn(err)
	}

	if db != nil {
		log.Println("Call DB success")
		return db, nil
	}

	db, err = createDB()

	if err != nil {
		return nil, errorConn(err)
	}

	log.Println("Call DB success")
	return db, nil
}

func createDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SUPER_USER"),
		os.Getenv("DB_SUPER_PASSWORD"),
		os.Getenv("DB_ROOT"))
	dbRoot, errRoot := gorm.Open(postgres.Open(config), &gorm.Config{})

	if errRoot != nil {
		return nil, errorConn(errRoot)
	}

	dbRoot.Exec(fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("DB_NAME")))

	sqlDbRoot, errRoot := dbRoot.DB()
	if errRoot != nil {
		return nil, errRoot
	}
	errRoot = sqlDbRoot.Close()
	if errRoot != nil {
		return nil, errRoot
	}

	config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, errorConn(err)
	}

	log.Println("create database success")
	return db, nil
}

func checkConn(db *gorm.DB) (*gorm.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errorConn(err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, errorConn(err)
	}

	log.Println("check database connection success")
	return db, nil
}

func migrateDB(db *gorm.DB) (*gorm.DB, error) {
	if err := db.AutoMigrate(
		&models.BranchTab{},
		&models.CustomerDataTab{},
		&models.LoanDataTab{},
		&models.MstCompanyTab{},
		&models.SkalaRentalTab{},
		&models.StagingCustomer{},
		&models.StagingError{},
		&models.VehicleDataTab{},
		&models.IdTab{},
		&models.User{}); err != nil {
		return nil, errorConn(err)
	}

	log.Println("migrate database success..")
	return db, nil
}

func seedInitialIdTab(db *gorm.DB) {
	initial := models.IdTab{
		CODE:  "006",
		DIGIT: 10,
		VALUE: 1}
	seedTable(db, &models.IdTab{}, initial)
}

func seedTable(db *gorm.DB, table any, newRecords any) {
	if !db.Migrator().HasTable(table) {
		return
	}

	if err := db.First(table).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		res := db.Create(newRecords)
		if res.Error != nil {
			log.Println(res.Error.Error())
		}
	}
}
