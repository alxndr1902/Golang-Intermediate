package main

import (
	"errors"
	"fmt"
)

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: empty data")
	}
	return nil
}
func main() {
	//data := []byte{}
	//if err := process(data); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println("Data processed successfully")

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("readConfig")
}
