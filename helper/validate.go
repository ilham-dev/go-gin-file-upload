package helper

import "mime/multipart"

func Validate(file *multipart.FileHeader, TypeFile string) bool {
	if file.Header.Get("Content-Type") != TypeFile {
		return false
	}else {
		return  true
	}
}
