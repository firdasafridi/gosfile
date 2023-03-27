# Gosfile: A File Encryption Library
`Gosfile` is a library that provides an easy-to-use solution for encrypting and decrypting files. Its primary goal is to enhance the security of files stored in storage by providing strong encryption.

## Features
Encrypt files using symmetric encryption.
Decrypt files that were encrypted using the library.
Easy-to-use API for file encryption and decryption.
Strong security and encryption protocols.
Installation
To install Gosfile, simply run:

```BASH
go get github.com/firdasafridi/gosfile
```

## Usage

Here's an example of how to use Gosfile to encrypt a file:

```go
package main

import (
	"fmt"
	"github.com/firdasafridi/gosfile"
)

func main() {
	key := []byte("16-byte-secret-key") // Your secret key should be 16, 24 or 32 bytes long

	// Encrypt file
	err := gosfile.EncryptFile("plain.txt", "encrypted.txt", key)
	if err != nil {
		panic(err)
	}

	fmt.Println("File encrypted successfully.")
}
```
And here's an example of how to decrypt the encrypted file:

```go
package main

import (
	"fmt"
	"github.com/firdasafridi/gosfile"
)

func main() {
	key := []byte("16-byte-secret-key") // Your secret key should be 16, 24 or 32 bytes long

	// Decrypt file
	err := gosfile.DecryptFile("encrypted.txt", "decrypted.txt", key)
	if err != nil {
		panic(err)
	}

	fmt.Println("File decrypted successfully.")
}
```

## Security
Gosfile uses strong encryption protocols to protect your files, including AES encryption in CBC mode with PKCS#7 padding. However, please note that security is only as strong as your secret key. Therefore, it is important to choose a strong key and keep it secure.

## Contributions
Contributions to Gosfile are welcome and encouraged! If you have any suggestions or ideas for improvement, feel free to submit a pull request or open an issue.

## Source sample
- [Base64 Output](https://github.com/firdasafridi/gosfile/tree/main/example/base64)
- [File Output](https://github.com/firdasafridi/gosfile/tree/main/example/file)
- [Sample HTTP](https://github.com/firdasafridi/gosfile/tree/main/example/http)
