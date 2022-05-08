package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type GetImageHandler struct {
	postClientAddress string
}

func NewGetImageHandler(postClientAddress string) Handler {
	return &GetImageHandler{
		postClientAddress: postClientAddress,
	}
}

func (handler *GetImageHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/posts/postImage/{imageName}", handler.GetPostImage)
	if err != nil {
		panic(err)
	}
}

func (handler *GetImageHandler) GetPostImage(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	imageName := pathParams["imageName"]
	fmt.Println(imageName)

	var base64Encoding string

	//check if image exists
	fileInfo, err := os.Stat("C:\\Users\\bogda\\OneDrive\\Desktop\\IV godina\\XML\\DislinktXWS-back\\temp-images\\" + imageName)
	if os.IsNotExist(err) {
		base64Encoding = ""
	} else {
		fmt.Println(fileInfo)
		// Read the entire file into a byte slice
		bytes, err := ioutil.ReadFile("C:\\Users\\bogda\\OneDrive\\Desktop\\IV godina\\XML\\DislinktXWS-back\\temp-images\\" + imageName)
		if err != nil {
			log.Fatal(err)
		}

		// Determine the content type of the image file
		mimeType := http.DetectContentType(bytes)

		// Prepend the appropriate URI scheme header depending
		// on the MIME type
		switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/png;base64,"
		}

		// Append the base64 encoded output
		base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	}

	response, err := json.Marshal(base64Encoding)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
