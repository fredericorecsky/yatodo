package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type DbConfig struct {
	Dbh *gorm.DB
	Err error
	database string
	dsn string
}

func (dbc *DbConfig) Setup(){
	backend, exists := os.LookupEnv("DB_BACKEND")
	if  ! exists  {
		fmt.Println("You need to define the DB_BACKEND with [sqllite|mysql]")
		os.Exit(1)
	}
	dbc.database = backend

	dsn, existsdsn := os.LookupEnv("DB_DSN")
	if ! existsdsn {
		fmt.Println("You need to define the DB_DSN with your database dsn details")
		os.Exit(1)
	}
	dbc.dsn = dsn
}

func (dbc *DbConfig) Connect() {

	dbc.Dbh, dbc.Err = gorm.Open( dbc.database, dbc.dsn)
	//dbc.Dbh, dbc.Err = gorm.Open("mysql", "yatodo:yatodosalt@/yatodo?charset=utf8&parseTime=True&loc=Local")
	if dbc.Err != nil {
		panic("Failed fo connect to database")
	}
}
