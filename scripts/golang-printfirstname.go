package main

import (
	"fmt" // for printing
	"io/ioutil" // for file reading
	"strings" // for string trim and split
	"os" // for command line args
)

func main(){
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an valid file name as input")
		return
	}
	fContent, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Printf("File Read Error: %v \n", err)
		return
	}
	
	// convert bytes buffer to string
	fContentStr := strings.TrimSpace(string(fContent))

	//handle empty file content error
	if len(fContentStr) == 0 {
		fmt.Println("Empty file provided. Please input a file with content as <firstname>:<lastname>")
	} else {
		nameParts := strings.Split(fContentStr, ":")
		
		// handle invalid format error
		if len(nameParts) == 1 {
			fmt.Println("Invalid format input. Please provide a file with content as <firstname>:<lastname>")
		} else {
			fmt.Println("\nFIRST NAME from GOLANG PROGRAM:" + nameParts[0])
		}
	}
}

