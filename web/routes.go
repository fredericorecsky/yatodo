package web

import (
	"github.com/fredericorecsky/yatodo/app"
	"github.com/fredericorecsky/yatodo/models/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Middleware

type AuthHeader struct {
	Token string `header:"Token"`
}

func LoadApp(yatodo *app.Todo) gin.HandlerFunc {
	return func(c *gin.Context) {
		yatodo.Connect()
		//defer
		c.Set("yatodo", yatodo)

		c.Next()
		yatodo.Db.Close()
	}
}
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := AuthHeader{}
		app := c.MustGet("yatodo").(*app.Todo)

		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		var user todo.User
		err := app.GetUserFromSecret(&user, h.Token)
		if err.Error != nil {
			c.AbortWithError(http.StatusForbidden, err.Error)
			return
		}

		c.Set("yatodouser", &user)
		c.Next()
	}
}
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		app := c.MustGet("yatodo").(*app.Todo)
		user := c.MustGet("yatodouser").(*todo.User)

		listID, _ := strconv.ParseUint(c.Param("listid"), 10, 64)
		var todoList todo.TodoList
		todoList.ID = uint(listID)
		err := app.GetTodoList(&todoList, *user)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "retrieving todolist"})
		}
		c.Set("todolist", &todoList)
		c.Next()
	}
}

// Controllers

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"doc": "doc",
	})
}

func userPOST(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)

	var user todo.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.CreateUser(&user)
	if err.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
	}
	c.JSON(http.StatusOK, user)
}

func todolistGET(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)
	user := c.MustGet("yatodouser").(*todo.User)

	if len(c.Param("listid")) > 0 {
		listID, _ := strconv.ParseUint(c.Param("listid"), 10, 64)
		var todoList todo.TodoList
		todoList.ID = uint(listID)
		app.GetTodoList(&todoList, *user)
		c.JSON(http.StatusOK, todoList)
		return
	} else {
		var todoLists []todo.TodoList
		app.GetTodoLists(&todoLists, *user)
		c.JSON(http.StatusOK, todoLists)
		return
	}
}
func todolistPOST(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)
	user := c.MustGet("yatodouser").(*todo.User)

	var json todo.TodoList
	h := AuthHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	list := todo.TodoList{Name: json.Name}

	app.CreateTodoList(&list, *user)
	c.JSON(http.StatusOK, list)
}

func itemGET(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)
	todoList := c.MustGet("todolist").(*todo.TodoList)

	itemID, _ := strconv.ParseUint(c.Param("itemid"), 10, 64)

	if itemID > 0 {
		var item todo.Item
		item.ID = uint(itemID)
		app.GetItem(&item, *todoList)

		c.JSON(http.StatusOK, item)
		return
	} else {

		var items []todo.Item
		app.GetItems(&items, *todoList)

		c.JSON(http.StatusOK, items)
		return
	}
}
func itemPOST(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)
	todoList := c.MustGet("todolist").(*todo.TodoList)

	var json todo.Item
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.TodoListID = todoList.ID
	app.CreateItem(&json)
	c.JSON(http.StatusOK, json)

}
func itemDELETE(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)
	todoList := c.MustGet("todolist").(*todo.TodoList)

	itemID, _ := strconv.ParseUint(c.Param("itemid"), 10, 64)
	var item todo.Item
	item.ID = uint(itemID)
	app.GetItem(&item, *todoList)
	err := app.DeleteItem(&item)
	if err.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, nil)

}

func itemDueToPOST(c *gin.Context) {
	type DueDate struct {
		DueTo string
	}

	app := c.MustGet("yatodo").(*app.Todo)

	var json DueDate

	itemId, _ := strconv.ParseUint(c.Param("itemid"), 10, 64)

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02"
	dueto, errDate := time.Parse(layout, json.DueTo)

	if errDate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errDate.Error()})
		return
	}
	item := todo.Item{DueTo: &dueto}
	item.ID = uint(itemId)

	app.AddDueTo(&item)

	c.JSON(http.StatusOK, item)
}
func itemDueToDELETE(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)

	itemId, _ := strconv.ParseUint(c.Param("itemid"), 10, 64)

	item := todo.Item{}
	item.ID = uint(itemId)

	app.DeleteDueTo(&item)
}

func labelPOST(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)

	var json todo.Label
	itemId, _ := strconv.ParseUint(c.Param("itemid"), 10, 64)

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := app.AddLabel(&json, uint(itemId))
	if err.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, json)
}
func labelDELETE(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)

	itemId, _ := strconv.ParseUint(c.Param("itemid"), 10, 64)
	labelId, _ := strconv.ParseUint(c.Param("labelid"), 10, 64)

	err := app.DeleteLabel(uint(itemId), uint(labelId))
	if err.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func commentPOST(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)

	var json todo.Comment
	itemId, _ := strconv.ParseUint(c.Param("itemid"), 10, 64)

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	app.AddComment(uint(itemId), json.Comment)

	c.JSON(http.StatusOK, gin.H{"ItemID": itemId, "Label": json.Comment})
}

func health(c *gin.Context) {
	app := c.MustGet("yatodo").(*app.Todo)

	err := app.Check()
	if err != nil {
		c.AbortWithError(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, gin.H{ "status":"alive"})
}
