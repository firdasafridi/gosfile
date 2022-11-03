package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/firdasafridi/gosfile"
)

var (
	gosInterface = &gosfile.GosFile{}
)

func init() {
	gosInterface, _ = gosfile.New(&gosfile.Config{
		TypeEncryption: gosfile.AES,
		Key:            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	})
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.

	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gosInterface.EncryptFromByte(fileBytes).ExportTo("./sample/http.enc"))
	fmt.Println(gosInterface.DecryptFromPath("./sample/http.enc").ExportTo("./sample/cuba.jpg"))

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	fmt.Println(http.ListenAndServe(":10002", nil))
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}
