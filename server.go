package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/device/locate", D_Locate)
	router.Run(":8080")
}

func D_Locate(c *gin.Context) {
	fmt.Println("get it!!!!!!!!!")
}
