package api

import (
	"github.com/gin-gonic/gin"
	"gotool/utils"
	"net/http"
)

func connactionService(router *gin.Engine) error {
	router.POST("/connaction", func(c *gin.Context) {
		
		server := c.PostForm("server")
		if server == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "缺少server"})
			return
		}
		user := c.PostForm("user")
		if user == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "缺少user"})
			return
		}
		password := c.PostForm("password")
		if password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "缺少password"})
			return
		}
		utils.Connaction(c, server, user, password)
		
	})
	return nil
}
