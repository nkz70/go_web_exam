package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"webxam/controllers"
)

type UserHandler interface {
	FetchUserList(c *gin.Context)
	FindUser(c *gin.Context)
}

type userHandler struct {
	UserController controllers.UserController
}

func NewUserHandler(uc controllers.UserController) UserHandler {
	return &userHandler{
		UserController: uc,
	}
}

func (uh *userHandler) FetchUserList(c *gin.Context) {
	var ur controllers.UserRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		log.Println(err)
	}

	users, err := uh.UserController.FetchUserList(&ur)
	if err != nil {

	}

	c.JSON(http.StatusOK, users)
}

func (uh *userHandler) FindUser(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {

	}

	users, err := uh.UserController.FindUser(id)
	if err != nil {

	}

	c.JSON(http.StatusOK, users)
}
