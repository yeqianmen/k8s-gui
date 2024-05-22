package main

import "github.com/gin-gonic/gin"

func res_string(c *gin.Context) {
	c.String(200, "ok")

}
func res_json(c *gin.Context) {
	type userjoson struct {
		Name     string `json:"name"`
		Id       int
		Age      int
		Password string `json:"password"`
	}
	jsonUser := userjoson{
		Name:     "wang",
		Id:       1,
		Age:      18,
		Password: "qwer",
	}
	c.JSON(200, jsonUser)
	c.JSON(200, gin.H{"username": "Alice", "age": 18})
}
func main() {
	r := gin.Default()
	r.GET("/text", res_string)
	r.GET("/json", res_json)
	r.Run(":17000")
}
