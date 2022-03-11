package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"webxam/controllers"
	"webxam/middleware"
)

type QuestionHandler interface {
	FetchQuestionList(c *gin.Context)
	FindQuestion(c *gin.Context)
	CreateQuestion(c *gin.Context)
	UpdateQuestion(c *gin.Context)
	DeleteQuestion(c *gin.Context)
}

type questionHandler struct {
	QuestionController controllers.QuestionController
}

func NewQuestionHandler(qc controllers.QuestionController) QuestionHandler {
	return &questionHandler{
		QuestionController: qc,
	}
}

func (qh *questionHandler) FetchQuestionList(c *gin.Context) {
	var qr controllers.QuestionRequest
	if err := c.ShouldBindQuery(&qr); err != nil {
		middleware.S.Error(err)
	}

	question, err := qh.QuestionController.FetchQuestionList(&qr)
	if err != nil {
		middleware.S.Error(err)
	}

	c.JSON(http.StatusOK, question)
}

func (qh *questionHandler) FindQuestion(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		middleware.S.Error(err)
	}

	question, err := qh.QuestionController.FindQuestion(id)
	if err != nil {
		middleware.S.Error(err)
	}

	c.JSON(http.StatusOK, question)
}

func (qh *questionHandler) CreateQuestion(c *gin.Context) {
	var qr controllers.QuestionRequest
	if err := c.ShouldBindJSON(&qr); err != nil {
		middleware.S.Error(err)
	}
	middleware.L.Info("create user", zap.Object("request", &qr))

	question, err := qh.QuestionController.CreateQuestion(&qr)
	if err != nil {
		middleware.S.Error(err)
	}

	c.JSON(http.StatusCreated, question)
}

func (qh *questionHandler) UpdateQuestion(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		middleware.S.Error(err)
	}

	var qr controllers.QuestionRequest
	if err := c.ShouldBindJSON(&qr); err != nil {
		middleware.S.Error(err)
	}
	middleware.L.Info("update user", zap.Object("request", qr))

	question, err := qh.QuestionController.UpdateQuestion(id, &qr)
	if err != nil {
		middleware.S.Error(err)
	}

	c.JSON(http.StatusOK, question)
}

func (qh *questionHandler) DeleteQuestion(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		middleware.S.Error(err)
	}

	if err := qh.QuestionController.DeleteQuestion(id); err != nil {
		middleware.S.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
