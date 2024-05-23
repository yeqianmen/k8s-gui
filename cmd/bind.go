package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func bindQuery(c *gin.Context) {
	var userinfo UserInfo
	err := c.ShouldBindQuery(&userinfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "you are wrong"})
		return
	}
	c.JSON(200, gin.H{"json": userinfo})
}
func bindJson(c *gin.Context) {
	var userinfo UserInfo

	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "you are wrong"})
		return
	}
	c.JSON(200, gin.H{"json": userinfo})
}
func main() {
	r := gin.Default()
	r.POST("/json", bindJson)
	r.POST("/query", bindQuery)

	r.Run(":8080")

}
