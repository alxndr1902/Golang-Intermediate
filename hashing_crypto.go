package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	//hash := sha512.Sum512([]byte(password))
	//
	//fmt.Println(password)
	//fmt.Println(hash)
	//fmt.Printf("SHA-256 Hash hex value: %x\n", hash)

	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original salt: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// hash the password with salt
	signUpHash := hashPassword(password, salt)

	// store the password and password in database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("Hash:", signUpHash)
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Haso of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	//verify
	//retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}
	loginHash := hashPassword(password, decodedSalt)

	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
