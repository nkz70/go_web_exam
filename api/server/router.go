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

	// user
	ur := persistence.NewUserPersistence(db.GetDB())
	uc := controllers.NewUserController(ur)
	un := handler.NewUserHandler(uc)

	// question
	qr := persistence.NewQuestionPersistence(db.GetDB())
	qc := controllers.NewQuestionController(qr)
	qn := handler.NewQuestionHandler(qc)

	v := router.Group(c.GetString("server.version"))
	{
		// user
		v.GET("/users", un.FetchUserList)
		v.GET("/user/:id", un.FindUser)
		v.POST("/user", un.CreateUser)
		v.PUT("/user/:id", un.UpdateUser)
		v.DELETE("/user/:id", un.DeleteUser)

		// question
		v.GET("/questions", qn.FetchQuestionList)
		v.GET("/question/:id", qn.FindQuestion)
		v.POST("/question", qn.CreateQuestion)
		v.PUT("/question/:id", qn.UpdateQuestion)
		v.DELETE("/question/:id", qn.DeleteQuestion)
	}

	return router
}
