package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read
	//reader := bufio.NewReader(strings.NewReader("Hello, bufio package!ee\n"))
	//
	//// read data byte by byte
	//data := make([]byte, 20)
	//
	//n, err := reader.Read(data)
	//if err != nil {
	//	fmt.Println("Error reading data:", err)
	//	return
	//}
	//fmt.Printf("Read %d bytes: %s\n", n, data[:n])
	//
	//line, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("Error reading line:", err)
	//	return
	//}
	//fmt.Println("Read string:", line)

	// Write
	writer := bufio.NewWriter(os.Stdout)

	// writing byte slice
	data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
	}

	// writing string
	str := "This is string\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
