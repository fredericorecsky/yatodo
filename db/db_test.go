package db

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db := new(DbConfig)
	db.Connect()
	defer db.Dbh.Close()

	if db.Err != nil {
		t.Errorf("Error connecting to the database")
	}

	err := db.Dbh.DB().Ping()
	if err != nil {
		t.Errorf("Error pinging to the database")
	}
}
