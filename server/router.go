package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"webexam/config"
	"webexam/controllers"
	"webexam/db"
	"webexam/db/persistence"
	"webexam/handler"
	"webexam/middleware"
)

func NewRouter() *gin.Engine {
	c := config.GetConfig()
	router := gin.New()
	router.Use(middleware.Logger)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"DELETE",
			"PUT",
		},
	}))

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
