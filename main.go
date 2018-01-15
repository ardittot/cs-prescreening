package main

import (
	"github.com/gin-gonic/gin"
)

func main () {
	auth,_ = InitToken()
	router := gin.Default()
	router.POST("/sicd", RequestSICD)
	router.POST("/dhn", RequestDHN)
	router.POST("/kemendagri", RequestKEMENDAGRI)
	router.Run("0.0.0.0:8000")	//Listen adn serve on 0.0.0.0:8000
}
