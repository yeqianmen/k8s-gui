package main

import "github.com/gin-gonic/gin"

func resString(c *gin.Context) {
	c.String(200, "ok")

}
func resJson(c *gin.Context) {
	type UserJson struct {
		Name     string `json:"name"`
		Id       int
		Age      int
		Password string `json:"password"`
	}
	JsonUser := UserJson{
		Name:     "wang",
		Id:       1,
		Age:      18,
		Password: "qwer",
	}
	c.JSON(200, JsonUser)
	c.JSON(200, gin.H{"username": "Alice", "age": 18})
}
func main() {
	r := gin.Default()
	r.GET("/text", resString)
	r.GET("/json", resJson)
	r.Run(":17000")
}
