# ismoilroziboyev/go-pkg Package

Welcome to the official repository for the ismoilroziboyev/go-pkg package! This package aims to simplify common tasks in Go development by providing a collection of utility methods related to validation, errors, hashing, and more.

## Overview

The ismoilroziboyev/go-pkg Package includes a set of commonly used methods that cover various aspects of software development, making it easier for Go developers to perform tasks such as validation, error handling, hashing, and struct mapping.

## Features

1. **Dominant Color:**
   - Method: `dominantColor.GetDominantColor(content []byte) (color.RGBA, error)`
   - This method serves to extract the dominant color from a JPEG or PNG image by analyzing the provided content. It leverages the capabilities of two external packages, namely `github.com/lucasb-eyer/go-colorful` and `github.com/nfnt/resize`. The go-colorful package is utilized for its effective color processing functionalities, while the resize package contributes to the efficient handling of image resizing operations. The combination of these packages enhances the accuracy and performance of the color extraction process, allowing for a robust solution in determining the dominant color within images.

Special thanks to the authors and maintainers of github.com/lucasb-eyer/go-colorful and github.com/nfnt/resize for their valuable contributions to the development of this method. 

2. **Errors:**
   - Error structure:
     ```go
     type Error struct {
	      i18nmsg struct{
               Uz string
	           Ru string
	           En string
	      msg     string
	      code    int
      }
     
   - Methods:
             `errors.NewInternalServerError(err error) error` - returns error wrapped with internal server error code and multilang error message 
             `errors.NewInternalServerErrorf(format string, args ...interface{}) error` - returns error wrapped with given formatted message, internal server error code and multilang error message 
             `errors.NewInternalServerErrorw(msg string, err error) error` - returns error wrapped with given message and error's message, internal server error code and multilang error message
   - Description: The package includes common errors with useful formatting methods to enhance error handling in your Go applications.

4. **Hash:**
   - Methods: 
   
   		`hash.HashMD5(s string) string` - 
This method provides a convenient way to generate an MD5 hash for a given input string. It utilizes the cryptographic hash function MD5, which produces a 128-bit hash value often represented as a 32-character hexadecimal number. While MD5 is widely used for non-security-critical purposes like checksums and data integrity verification, it is considered weak for cryptographic security due to vulnerabilities. It is recommended to use stronger hash functions for security-sensitive applications.


		`hash.HashSHA1(s string) string` - 
This method implements the SHA-1 hashing algorithm to create a hash value for the provided input string. SHA-1 produces a 160-bit hash, typically represented as a 40-character hexadecimal string. However, SHA-1 is no longer considered secure against well-funded attackers, and its use is deprecated in cryptographic applications. For better security, consider using SHA-256 or SHA-3.

		`hash.HashSHA256(s string) string` - 
This function uses the SHA-256 hash function to generate a hash for the input string. SHA-256 is part of the SHA-2 family and produces a 256-bit hash, commonly represented as a 64-character hexadecimal string. It is widely used and considered secure for various cryptographic purposes, including digital signatures and data integrity verification.

		`hash.HashSHA512(s string) string` - 
The `HashSHA512` method employs the SHA-512 hashing algorithm to create a hash value for the provided string. SHA-512 produces a 512-bit hash, represented as a 128-character hexadecimal string. It offers a higher level of security due to its longer hash length. SHA-512 is suitable for applications requiring a robust hash function, such as cryptographic protocols and password hashing.

1. **Mapper:**
   - Method: `Map(ctx context.Context, input any, output any) error`
   - Description: The mapper package includes functions for mapping between two structs using their JSON formats (Output must be a pointer).

2. **Validation:**
   - Email Validation:
     - Method: `validation.IsValidEmail(email string) bool`
     - Description: Checks the validity of an email address using the net/mail package.
   - Phone Number Validation:
     - Method: `validation.IsValidPhoneNumber(phoneNumber string) bool`
     - Description: Checks the validity of a phone number in the '998900000000' format.

## Installation

To use the Your Go Utils Package in your Go project, you can run the following command:

```bash
go get -u github.com/ismoilroziboyev/go-pkg
```

## Examples

  1. **Dominant color**
   ```go
package main

import (
	"fmt"
	"github.com/ismoilroziboyev/go-pkg/dominantColor"
)

func main() {
	content := []byte{} // Replace this with actual image content

	dominantColor, err := dominantColor.GetDominantColor(content)
	if err != nil {
		fmt.Println("Error extracting dominant color:", err)
		return
	}

	fmt.Printf("Dominant Color: %#v\n", dominantColor)
}

```
2. **Errors**
```go
package main

import (
	"fmt"

	"github.com/ismoilroziboyev/go-pkg/errors"
)

func main() {
	// Example 1: Creating a new internal server error
	err1 := errors.NewInternalServerError(fmt.Errorf("database connection failed"))

	myErr1 := err1.(errors.Error)

	fmt.Println(myErr1.Code())             // prints "500"
	fmt.Println(myErr1.Error())            // prints "database connection failed"
	fmt.Println(myErr1.ErrorI18nMsg("uz")) //prints "Serverda xatolik"

	// Example 2: Creating a new internal server error with a formatted message
	err2 := errors.NewInternalServerErrorf("Failed to process request: %s", "invalid input")

	myErr2 := err2.(errors.Error)

	fmt.Println(myErr2.Code())             // prints "500"
	fmt.Println(myErr2.Error())            // prints "Failed to process request: invalid input"
	fmt.Println(myErr2.ErrorI18nMsg("en")) // prints "Internal server error"

	// Example 3: Creating a new internal server error with a custom message and wrapping an existing error
	err3 := errors.NewInternalServerErrorw("Failed to fetch data", fmt.Errorf("database query failed"))

	myErr3 := err3.(errors.Error)

	fmt.Println(myErr3.Code())             // prints "500"
	fmt.Println(myErr3.Error())            // prints "Failed to fetch data: database query failed"
	fmt.Println(myErr3.ErrorI18nMsg("uz")) // prints "Serverda xatolik"

	// Example 4: Wrapping error with extra context message
	err4 := errors.Wrap("Cannot send ping message", fmt.Errorf("connection timeout"))

	myErr4 := err4.(errors.Error)

	fmt.Println(myErr4.Code())             // prints "0"; wrapping error and wrapped error codes will be the same code
	fmt.Println(myErr4.Error())            // prints "Cannot send ping message: connection timeout"
	fmt.Println(myErr4.ErrorI18nMsg("uz")) // prints "Serverda xatolik"

	// Example 5: Checking if the error is wrapped by any error
	wrappingErr := fmt.Errorf("connection timeout")
	err5 := errors.Wrap("Cannot send ping message", wrappingErr)

	myErr5 := err5.(errors.Error)

	fmt.Println(errors.IsWrappedWith(myErr5, wrappingErr))                         //prints "true"
	fmt.Println(errors.IsWrappedWith(myErr5, fmt.Errorf("this is another error"))) //prints "false"
}

```
  3. **Hash**
```go
package main

import (
	"fmt"

	"github.com/ismoilroziboyev/go-pkg/hash"
)

func main() {
	hashingString := "this is the string"

	hashedString := hash.HashMD5(hashingString)

	fmt.Println(hashedString) //prints 'fa77937febc2a2a9754d326bb88e0b16'
}

```
  4. **Mapper**
```go
package main

import (
	"context"
	"fmt"

	"github.com/ismoilroziboyev/go-pkg/mapper"
)

type Person struct {
	Firstname  string `json:"firstname"`
	SecondName string `json:"secondname"`
	LastName   string `json:"lastname"`
}

type PartialPerson struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {

	person := Person{
		Firstname:  "First name",
		SecondName: "Second name",
		LastName:   "Last name",
	}

	var partialPerson PartialPerson

	if err := mapper.Map(context.Background(), person, &partialPerson); err != nil {
		fmt.Println("error occured while mapping person:", err.Error())
	}

	fmt.Println(partialPerson.Firstname) // prints "First name"
	fmt.Println(partialPerson.LastName)  // prints "Last name"
}

```
  5. **Validation**
```go
package main

import (
	"fmt"

	"github.com/ismoilroziboyev/go-pkg/validation"
)

func main() {
	//Example 1: Email validation
	validEmail := "email@example.com"

	fmt.Println(validation.IsValidEmail(validEmail)) // prints: "true"

	invalidEmail := "email.com"

	fmt.Println(validation.IsValidEmail(invalidEmail)) //prints: "false"

	//Example 2: Phone number validation
	validPhone := "998939444321"

	fmt.Println(validation.IsValidPhone(validPhone)) //prints: "true"

	invalidPhone := "9444321"

	fmt.Println(validation.IsValidPhone(invalidPhone)) //prints: "false"
}

```

## Contact
If you have any questions or feedback, feel free to reach out to Ismoil Ro'ziboyev via email at [roziboyevismoil3@gmail.com].


Happy coding!
