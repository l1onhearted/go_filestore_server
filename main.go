package main

import (
	"fmt"
	"go_filestore_server/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandeler)
	http.HandleFunc("/file/upload/suc", handler.UploadSuccess)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed to start,err:", err.Error())
	}
}
