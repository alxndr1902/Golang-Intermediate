package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type customError struct {
	code    int
	message string
	er      error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s %v\n", e.code, e.message, e.er)
}

//func doSomething() error {
//	return &customError{
//		code:    500,
//		message: "Something went wrong",
//	}
//}

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "something went wrong!",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal error")
}
