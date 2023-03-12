package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const passwordLength = 12 // Change this to adjust the password length

func main() {
	// Generate a random byte slice
	randBytes := make([]byte, passwordLength)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}

	// Encode the byte slice as a base64 string
	password := base64.StdEncoding.EncodeToString(randBytes)[:passwordLength]

	// Print the generated password
	fmt.Println(password)
}
