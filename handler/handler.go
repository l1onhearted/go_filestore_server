package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandeler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //返回上传html页面
		data, err := ioutil.ReadFile("/home/zhangyuankun/go/src/filestore_server/handler/static/view/index.html")
		if err != nil {
			io.WriteString(w, "iner server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" { //接受文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Failed to get error:", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("./tmpSaved/" + head.Filename)
		if err != nil {
			fmt.Printf("failed to create file,err:%s\n", err.Error())
			return
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("failed to save data into file,err:%s\n", err.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSuccess(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Success")
}
