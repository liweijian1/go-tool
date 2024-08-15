package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func CreateApi() {
	router := gin.Default()
	if err := connactionService(router);err!=nil {
		log.Fatalf("Failed to configure service: %+v", err)
	}
	router.Run(":8080")
}
