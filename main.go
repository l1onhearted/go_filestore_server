package main

import (
	"filestore_server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandeler)
	http.HandleFunc("/file/upload/suc", handler.UploadSuccess)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed to start,err:", err.Error())
	}
}
