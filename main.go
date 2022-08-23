package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"mime"
	"net/http"
)

var message string

func handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	message = ""
	if r.Method == "POST" {
		// Parse files
		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			message = "No File Selected"
			tmpl.Execute(w, message)
			return
		}
		defer file.Close()
		// Get File extension
		ext, _ := mime.ExtensionsByType(handler.Header.Get("Content-Type"))
		// Create a Temporary File
		patern := fmt.Sprintf("upload-*%s", ext[0])
		tempFile, err := ioutil.TempFile("uploads", patern)
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		tempFile.Write(fileBytes)

		if err != nil {
			fmt.Println(err)
		}
		message = "File Upload Successful"
	}

	tmpl.Execute(w, message)
}
func main() {
	uploadFs := http.FileServer(http.Dir("uploads"))
	staticFs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", handler)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", uploadFs))
	http.Handle("/static/", http.StripPrefix("/static/", staticFs))
	fmt.Println("[INFO] Server Running on http://127.0.0.1:8888")
	http.ListenAndServe(":8888", nil)
}
