package todo

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique_index;not null"`
	SecretToken string `gorm: "unique;not null"`
	TodoLists   []TodoList
}

func (user *User) UpdateToken() {
	h := sha256.Sum256([]byte(user.Username))
	user.SecretToken = hex.EncodeToString(h[:])
}

type TodoList struct {
	gorm.Model
	Name   string
	UserID uint
	Items  []Item
}

type Item struct {
	gorm.Model
	Description string
	DueTo       *time.Time
	DoneAt      *time.Time
	TodoListID  uint
	Labels      []*Label `gorm:"many2many:items_labels;"`
	Comments    []Comment
}

type Label struct {
	gorm.Model
	Label string
	Items []*Item `gorm:"many2many:items_labels;"`
}

type Comment struct {
	gorm.Model
	Comment string
	ItemID  uint
}
