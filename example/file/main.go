package main

import (
	"fmt"

	"github.com/firdasafridi/gosfile"
)

func main() {
	gosInterface, _ := gosfile.New(&gosfile.Config{
		TypeEncryption: gosfile.AES,
		Key:            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	})

	err := gosInterface.EncryptFromPath("./sample/sample.jpg").ExportTo("./sample/output.enc")
	fmt.Println(err)

	err = gosInterface.DecryptFromPath("./sample/output.enc").ExportTo("./sample/plain.jpg")
	fmt.Println(err)
}
