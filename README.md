# gosfile: File Encryption Library
`gosfile` is a robust library for file encryption, designed with the goal of securing files stored in storage. It offers a straightforward way to encrypt and decrypt files, ensuring your data remains inaccessible and secure.

## Features
- Supports encryption and decryption of files.
- Allows encryption from various input types - paths, bytes, and base64 strings.
- Uses gocrypt for the underlying encryption and decryption process. (this support multiple encryption)

## Installation
You can install gosfile using go get:
```sh
go get -u github.com/firdasafridi/gosfile
```

## Usage
Here's a basic example of how to use gosfile to encrypt and decrypt a file.

```go
package main

import (
	"fmt"
	"github.com/firdasafridi/gosfile"
)

func main() {
	config := &gosfile.Config{
		TypeEncryption: gosfile.AES,
		Key:            "your-encryption-key",
	}

	gosInterface, err := gosfile.New(config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = gosInterface.EncryptFromPath("./sample/sample.jpg").ExportTo("./sample/output.enc")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
### Decryption
```go
package main

import (
	"fmt"
	"github.com/firdasafridi/gosfile"
)

func main() {
	// ... (assuming gosInterface is already defined and initialized)

	err = gosInterface.DecryptFromPath("./sample/output.enc").ExportTo("./sample/plain.jpg")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
Replace `"your-encryption-key"` with your actual encryption key.

## Reference
### New
```go
func New(cfg *Config) (*GosFile, error)
```
Creates a new GosFile instance with the provided configuration.

### EncryptFromPath
```go
func (gosFile *GosFile) EncryptFromPath(path string) *Option
```

Encrypts a file from the given path.

### DecryptFromPath
```go
func (gosFile *GosFile) DecryptFromPath(path string) *Option
```

Decrypts a file from the given path.

## Supported MIME Types
The library supports various MIME types. They can be configured in the Config struct passed to the New function.

## Examples
Check out more detailed examples:

- [Base64 Output](https://github.com/firdasafridi/gosfile/tree/main/example/base64)
- [File Output](https://github.com/firdasafridi/gosfile/tree/main/example/file)
- [Sample HTTP](https://github.com/firdasafridi/gosfile/tree/main/example/http)
