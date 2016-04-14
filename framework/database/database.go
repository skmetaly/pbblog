package database

import (
	"database/sql"
	//"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	//Postgres lib
	_ "github.com/lib/pq"

	"log"
)

//DatabaseConfig is used for storing database credentials config
type DatabaseConfig struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`
	Host     string `json:"Host"`
}

//Database type
type Database struct {
	Config        DatabaseConfig
	Connection    sql.DB
	ORMConnection gorm.DB
}

//Connect opens a new connection for SQL DB
func (d *Database) Connect() error {
	db, err := sql.Open("postgres", "postgres://"+d.Config.Username+"@"+d.Config.Host+"/"+d.Config.Database+"?sslmode=disable")

	if err != nil {
		log.Fatalln("Could not connect to database", err)
		panic(err)
	}

	d.Connection = *db

	return err
}

//ConnectORM opens a new ORM connection to the database
func (d *Database) ConnectORM() error {

	Db, err := gorm.Open("postgres", "user="+d.Config.Username+" dbname="+d.Config.Database+" sslmode=disable")

	if err != nil {
		log.Fatalln("Could not connect to database from ORM", err)
		panic(err)
	}

	d.ORMConnection = *Db

	return err
}

//SetConfig is a setter for DatabaseConfig
func (d *Database) SetConfig(conf DatabaseConfig) {
	d.Config = conf
}
