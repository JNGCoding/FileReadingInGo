/*
Program Briefing :
This project scans a binary file and prints it's content to the standard output.

Program flow :
1) Open the file.
2) Read it's data
3) Print.
*/

package main

import (
	"fmt"
	"log"
	"os"
)

// Program entry point
func main() {
	ReadFile, err := os.Open(os.Args[1]) // os.Args stores all the program arguments.
	if err != nil {
		log.Fatal("Unable to Open.")
	}

	// defer keyword runs a line of code right before exiting the code.
	// It follows the execution flow of First-In ---> Last-Out
	defer ReadFile.Close()

	var data []byte = []byte{0} // An array of any type is declared this way. You can also use "make()" to allocate heap memory or create a pointer.
	for {
		bytes_read, err := ReadFile.Read(data) // Read len(data) bytes. bytes_read will be the amount of bytes read.
		fmt.Print(data[0], " ")                // Print is pretty much same as the print() in Python.
		if err != nil || bytes_read == 0 {
			break
		}
	}
	fmt.Println("")

	fmt.Println("Program Done.")
}
