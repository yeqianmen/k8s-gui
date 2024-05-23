package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func query(c *gin.Context) {
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.Query("user"))
	fmt.Println(c.QueryArray("user"))
}
func main() {
	r := gin.Default()
	r.GET("/query/:userId/:bookId", query)
	r.Run(":8080")
}
