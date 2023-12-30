package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"api/src/config/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Inst *DataSourceInst
var dsc *DataSource

type DataSourceInst struct {
	*sqlx.DB
}

type DataSource struct {
	Driver      string
	Host        string
	Port        string
	User        string
	Pass        string
	Database    string
	SSLmode     string
	Options     string
	PoolMax     int
	PoolMin     int
	PoolAcquire int
	PoolIdle    int
}

func (ds *DataSource) New() *DataSource {
	return &DataSource{
		Driver:   "postgres",
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Pass:     os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		SSLmode:  os.Getenv("POSTGRES_SSL_MODE"),
	}
}

func (ds *DataSource) Handshake() string {
	handshake := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		ds.Host,
		ds.Port,
		ds.User,
		ds.Pass,
		ds.Database,
		ds.SSLmode,
	)
	return handshake
}

func (ds *DataSource) Conn() error {
	ds = ds.New()
	dsn := ds.Handshake()
	var err error

	db, err := sqlx.Open(ds.Driver, dsn)
	if err != nil {
		time.Sleep(10 * time.Second)
		db, err = sqlx.Open(ds.Driver, dsn)
		if err != nil {
			logger.Error("[DATABASE CONNECTION ERR]", err)
			return err
		}
	}

	Inst = &DataSourceInst{DB: db}
	return nil
}

func (ds *DataSource) Close() error {
	if err := Inst.Close(); err != nil {
		logger.Error("[ERROR CLOSED DB]", err)
		return errors.New("Erro interno.")
	}

	return nil
}

func (d *DataSource) TestConn(db *sqlx.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return nil
}
