package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"webxam/config"
)

var (
	db *gorm.DB
)

type dbConf struct {
	Db
}

type Db struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
	Dsn      string
}

func Init() {
	var dbc dbConf
	c := config.GetConfig()
	if err := c.Unmarshal(&dbc); err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf(dbc.Dsn, dbc.Host, dbc.User, dbc.Password, dbc.Dbname, dbc.Port)
	log.Println(dsn)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connect succeed")
}

func GetDB() *gorm.DB {
	return db
}
