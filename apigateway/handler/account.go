package handler

import (
	"apigw/middleware"
	proto "apigw/proto/account"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	AccountServiceClient proto.AccountService
)

func RegisterHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	res, err := AccountServiceClient.AccountRegister(middleware.GetTracerContext(c), &proto.ReqAccountRegister{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -2,
			"message": "server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    res.Code,
		"message": res.Message,
	})
	return
}
