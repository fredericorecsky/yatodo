package web

import (
	"github.com/fredericorecsky/yatodo/app"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {

	yatodo := app.Todo{}

	r := gin.New()

	//r.Use(gin.Recovery())
	r.Use(LoadApp(&yatodo))
	r.GET("/", index)
	r.POST("/users", userPOST)

	r.GET("/health", health)
	r.GET("/ready", health)

	authenticated := r.Group("/todolists")
	authenticated.Use(Authentication())

	authenticated.GET("", todolistGET)
	authenticated.POST("", todolistPOST)

	authorizated := authenticated.Group("/")
	authorizated.Use(Authorization())

	authorizated.GET("/:listid", todolistGET)

	authorizated.POST("/:listid/items", itemPOST)
	authorizated.GET("/:listid/items", itemGET)
	authorizated.GET("/:listid/items/:itemid", itemGET)
	authorizated.DELETE("/:listid/items/:itemid", itemDELETE)

	authorizated.POST("/:listid/items/:itemid/labels", labelPOST)
	authorizated.DELETE("/:listid/items/:itemid/labels/:labelid", labelDELETE)

	authorizated.POST("/:listid/items/:itemid/comments", commentPOST)
	authorizated.POST("/:listid/items/:itemid/dueto", itemDueToPOST)
	authorizated.DELETE("/:listid/items/:itemid/dueto", itemDueToDELETE)

	return r
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	server := SetRouter()

	err := server.Run(":9000")

	_ = err
}
