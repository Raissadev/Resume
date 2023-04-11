package config

import (
	. "api/src/utils"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var env = Lenv.New()
var Instance *gorm.DB

type DataSource struct {
	Host     string
	Port     string
	Password string
	User     string
	Database string
	SSLmode  string
}

type DataSourceInterface interface {
	New() *DataSource
}

func (ds *DataSource) New() *DataSource {
	return &DataSource{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		SSLmode:  "disable",
	}
}

func (ds *DataSource) GetDSN() string {
	s := ds.New()

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", s.Host, s.User, s.Password, s.Database, s.Port, s.SSLmode)
}

func (ds *DataSource) Open() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  ds.GetDSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	dbs, err := db.DB()

	dbs.SetConnMaxIdleTime(time.Minute * 5)
	dbs.SetMaxIdleConns(10)
	dbs.SetMaxOpenConns(100)
	dbs.SetConnMaxLifetime(time.Hour)

	db.Set("gorm:insert_modifier", "IGNORE")

	Instance = db

	return err
}

func CloseDB() error {
	db, _ := Instance.DB()

	return db.Close()
}
