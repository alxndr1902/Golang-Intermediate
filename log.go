package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hello from Log")
	log.SetPrefix("INFO: ")
	log.Println("Hello from log with prefix")

	// log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Hello from log with flags")
	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file:", err)
	}
	defer file.Close()

	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is an debug message")
	warnLogger1.Println("This is an warning message")
	infoLogger1.Println("This is an info message")
	errorLogger1.Println("This is an error message")
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
