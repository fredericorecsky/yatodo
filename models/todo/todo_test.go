package todo

import (
	"github.com/fredericorecsky/yatodo/db"
	"testing"
)

func TestConnect(t *testing.T) {
	connection := new(db.DbConfig)

	connection.Connect()
}

func TestUser_UpdateToken(t *testing.T) {
	user := User{Username: "Huub"}
	user.UpdateToken()

	if len(user.SecretToken) == 0 {
		t.Errorf("Update token did not updated the token")
	}
}
