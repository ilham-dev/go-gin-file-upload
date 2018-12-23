### :chicken: Go Gin File Upload Example 
    I just Learning GO Programming With Uploading File
### Docs
   [Gin upload file example](https://github.com/gin-gonic/gin/tree/master/examples/upload-file/multiple)
    
### Install GIN
    go get github.com/gin-gonic/gin
    
### Filter File Type
    if file.Header.Get("Content-Type") != "image/png" {
    	c.JSON(http.StatusOK, JsonStatus{400,"Not Image"})
    	return
    }
    
### Run
    go run main.go
    
## Happy Code :computer: