package app

import (
	"github.com/fredericorecsky/yatodo/db"
	"github.com/fredericorecsky/yatodo/models/todo"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Db *gorm.DB
}

func (app *Todo) Connect() {
	db := new(db.DbConfig)
	db.Setup()
	db.Connect()
	app.Db = db.Dbh

}

func (app *Todo) Migrate() {
	app.Connect()
	defer app.Db.Close()

	app.Db.AutoMigrate(&todo.User{})
	app.Db.AutoMigrate(&todo.TodoList{})
	app.Db.AutoMigrate(&todo.Item{})
	app.Db.AutoMigrate(&todo.Label{})
	app.Db.AutoMigrate(&todo.Comment{})
}

func (app *Todo) Check() error {
	err := app.Db.DB().Ping()
	return err

}

func (app *Todo) CreateUser(user *todo.User) *gorm.DB {
	user.UpdateToken()
	err := app.Db.Create(&user)
	return err
}

func (app *Todo) GetUserFromSecret(user *todo.User, token string) *gorm.DB {
	err := app.Db.Where("secret_token = ?", token).First(&user)
	return err
}

func (app *Todo) CreateTodoList(list *todo.TodoList, user todo.User) *gorm.Association {
	err := app.Db.Model(&user).Association("TodoLists").Append(list)
	return err
}

func (app *Todo) GetTodoLists(todolists *[]todo.TodoList, user todo.User) *gorm.DB {
	err := app.Db.Where("user_id = ?", user.ID).Find(&todolists)
	return err
}

func (app *Todo) GetTodoList(todolist *todo.TodoList, user todo.User) *gorm.DB {
	err := app.Db.Where("user_id = ? and id = ?", user.ID, todolist.ID).Find(&todolist)
	return err
}

func (app *Todo) GetItems(items *[]todo.Item, todolist todo.TodoList) *gorm.DB {
	//err := app.Db.Where( "todo_list_id = ?", todolist.ID).Find( &items )
	err := app.Db.Preload("Comments").Preload("Labels").Where("todo_list_id = ?", todolist.ID).Find(&items)
	return err
}

func (app *Todo) GetItem(item *todo.Item, todolist todo.TodoList) *gorm.DB {
	err := app.Db.Preload("Comments").Preload("Labels").Where("todo_list_id = ? and id = ?", todolist.ID, item.ID).Find(&item)
	return err
}

func (app *Todo) CreateItem(item *todo.Item) *gorm.Association {
	var list todo.TodoList
	app.Db.Where("ID = ?", item.TodoListID).First(&list)
	err := app.Db.Model(&list).Association("Items").Append(item)
	return err
}

func (app *Todo) DeleteItem(item *todo.Item) *gorm.DB {
	err := app.Db.Delete(&item)
	return err
}

func (app *Todo) AddLabel(label *todo.Label, itemId uint) *gorm.Association {
	var item todo.Item
	app.Db.Where("id = ?", itemId).First(&item)

	app.Db.FirstOrCreate(&label, todo.Label{Label: label.Label})

	err := app.Db.Model(&item).Association("Labels").Append(label)
	return err
}

func (app *Todo) DeleteLabel(itemId uint, labelId uint) *gorm.Association {
	var item todo.Item
	app.Db.Where("id = ?", itemId).First(&item)

	var label todo.Label
	app.Db.Where("id = ?", labelId).First(&label)

	err := app.Db.Model(&item).Association("Labels").Delete(&label)
	return err
}

func (app *Todo) AddComment(itemId uint, value string) *gorm.Association {
	// ITEM NOT FOUND
	var item todo.Item
	app.Db.Where("id = ?", itemId).First(&item)

	comment := todo.Comment{Comment: value}

	err := app.Db.Model(&item).Association("Comments").Append(comment)
	return err
}

func (app *Todo) AddDueTo(item *todo.Item) *gorm.DB {
	dueto := item.DueTo
	app.Db.Where("id = ?", item.ID).First(&item)
	err := app.Db.Model(&item).Update("DueTo", dueto)

	return err
}

func (app *Todo) DeleteDueTo(item *todo.Item) *gorm.DB {
	err := app.Db.Model(&item).Update("DueTo", nil)
	return err
}
