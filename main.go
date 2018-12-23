package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-file-upload/helper"
)

func main()  {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", helper.Upload)
	_ = router.Run()
}
