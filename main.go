package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	output := ""
	byteFlag := flag.Bool("c", false, "Output the number of bytes in a file.")
	lineFlag := flag.Bool("l", false, "Output the number of lines in a file.")
	wordFlag := flag.Bool("w", false, "Output the number of words in a file.")
	charFlag := flag.Bool("m", false, "Output the number of characters in a file.")
	
	flag.Parse()

	args := flag.Args()

	var fileData []byte

	if len(args) > 0 {
		// Read from file
		file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fileData = file
	} else {
		// Read from standard input
		stdin, err := io.ReadAll(os.Stdin)
		
		if err != nil {
			panic(err)
		}

		fileData = stdin
	}

	switch {
		case *charFlag: 
			output += fmt.Sprintf("%d ", CalculateNoOfChars(fileData))

		case *lineFlag:
			output += fmt.Sprintf("%d ", CalculateNoOfLines(fileData))

		case * wordFlag:
			output += fmt.Sprintf("%d ", CalculateNoOfWords(fileData))

		case *byteFlag:
			output += fmt.Sprintf("%d ", CalculateNoOfBytes(fileData))

		default:
			output += fmt.Sprintf("%d %d %d ", CalculateNoOfLines(fileData), CalculateNoOfWords(fileData), CalculateNoOfBytes(fileData))
	}

	if len(args) > 0 {
		output += args[0]
	}

	fmt.Println(output)
}

