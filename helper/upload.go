package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type JsonStatus struct {
	Status int
	Message string
}

func Validate(c *gin.Context)  {

}


func Upload(c *gin.Context)  {
	//name := c.PostForm("name")
	//email := c.PostForm("email")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]


	for _, file := range files {
		// Validate Your File Type
		if file.Header.Get("Content-Type") != "image/png" {
			c.JSON(http.StatusOK, JsonStatus{400,"Not PNG"})
			return
		}

		filename := filepath.Base(file.Filename)

		if err := c.SaveUploadedFile(file, "./files/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, JsonStatus{200,"Berhasil"})
}