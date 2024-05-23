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
func form(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	s := c.PostFormArray("name")
	fmt.Printf("s: %v\n", s)
	c.JSON(200, gin.H{
		"nameArray": s,
		"name":      name,
	})
	city := c.DefaultPostForm("city", "JiNan")
	c.JSON(200, gin.H{
		"city": city,
	})
	fmt.Println(c.MultipartForm())
}

func main() {
	r := gin.Default()
	r.GET("/query/:userId/:bookId", query)
	r.POST("/list", form)
	r.Run(":8080")
}
