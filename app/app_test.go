package app

import (
	"github.com/fredericorecsky/yatodo/db"
	"github.com/fredericorecsky/yatodo/models/todo"
	"testing"
)

func TestMigrate(t *testing.T) {

}

func TestBootstrap(t *testing.T) {

	db := new(db.DbConfig)
	db.Connect()
	defer db.Dbh.Close()

	app := Todo{Db: db.Dbh}

	app.migrate()

	user := app.CreateUser("witcher")

	listitem := app.CreateTodoList(user, "Test list")

	item1 := app.CreateItem(listitem, todo.Item{Description: "Finish docker"})
	item2 := app.CreateItem(listitem, todo.Item{Description: "Finish k8s"})

	label := app.AddLabel(item1, todo.Label{Label: "Must finish!"})
	label = app.AddLabel(item2, label)

	comment := app.AddComment(item1, todo.Comment{Comment: "First run local"})
	comment2 := app.AddComment(item2, todo.Comment{Comment: "After on k8s"})

	_ = comment
	_ = comment2

}
