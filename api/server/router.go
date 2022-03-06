package server

import (
	"github.com/gin-gonic/gin"

	"webxam/config"
	"webxam/controllers"
	"webxam/db"
	"webxam/db/persistence"
	"webxam/handler"
	"webxam/middleware"
)

func NewRouter() *gin.Engine {
	c := config.GetConfig()
	router := gin.New()
	router.Use(middleware.Logger)

	ur := persistence.NewUserPersistence(db.GetDB())
	uc := controllers.NewUserController(ur)
	un := handler.NewUserHandler(uc)

	v := router.Group(c.GetString("server.version"))
	{
		v.GET("/users", un.FetchUserList)
		v.GET("/user/:id", un.FindUser)
	}

	return router
}
