package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"webexam/controllers"
	"webexam/middleware"
)

type UserHandler interface {
	FetchUserList(c *gin.Context)
	FindUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
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
		middleware.S.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	users, err := uh.UserController.FetchUserList(&ur)
	if err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *userHandler) FindUser(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	users, err := uh.UserController.FindUser(id)
	if err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *userHandler) CreateUser(c *gin.Context) {
	var ur controllers.UserRequest
	if err := c.ShouldBindJSON(&ur); err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	middleware.L.Info("create user", zap.Object("request", ur))

	user, err := uh.UserController.CreateUser(&ur)
	if err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uh *userHandler) UpdateUser(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var ur controllers.UserRequest
	if err := c.ShouldBindJSON(&ur); err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	middleware.L.Info("update user", zap.Object("request", ur))

	user, err := uh.UserController.UpdateUser(id, &ur)
	if err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *userHandler) DeleteUser(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := uh.UserController.DeleteUser(id); err != nil {
		middleware.S.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
