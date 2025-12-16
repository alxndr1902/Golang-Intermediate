package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, base64 encoding")

	// encode to Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded:", encoded)

	// decode from base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))

	// URL safe,avoid '/' and '+'

	urlSafeEncoded := base64.URLEncoding.EncodeToString(decoded)
	fmt.Println("URLSafeEncoded:", urlSafeEncoded)
}
