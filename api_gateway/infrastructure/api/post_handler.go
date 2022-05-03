package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type PostHandler struct {
	postClientAddress string
}

func NewPostHandler(postClientAddress string) Handler {
	return &PostHandler{
		postClientAddress: postClientAddress,
	}
}

func (handler *PostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/posts/postImage", handler.InsertPost)
	if err != nil {
		panic(err)
	}
}

func (handler *PostHandler) InsertPost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	uploadFile(w, r)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("Image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("D:\\DislinktXWS-back\\temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
}
