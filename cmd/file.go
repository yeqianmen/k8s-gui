package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func upLoadFile(c *gin.Context) {
	f, err := c.FormFile("file")
	//log.Println(f.Filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.SaveUploadedFile(f, "./"+f.Filename)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	}
}
func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 //memory limit
	r.POST("/upload", upLoadFile)
	r.Run(":8080")
}
