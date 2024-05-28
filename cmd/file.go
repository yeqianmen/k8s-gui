package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func upLoadFile(c *gin.Context) {
	f, err := c.FormFile("file") // form-name is file,matches curl
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
func uploadMuti(c *gin.Context) {
	form, _ := c.MultipartForm()

	/*	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {

	*/
	files := form.File["files"] // form name is files,matches curl
	fmt.Println(files)
	for _, file := range files {

		if err := c.SaveUploadedFile(file, "./"+file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}
	c.String(200, fmt.Sprintf("%d files uploaded", len(files)))
}

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 //memory limit
	r.POST("/upload", upLoadFile)
	r.POST("/uploadmuti", uploadMuti)
	r.Run(":8080")
}
