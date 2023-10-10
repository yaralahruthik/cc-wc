package main

import "bytes"

func CalculateNoOfBytes(content []byte) int {
	return len(content)
}

func CalculateNoOfLines(content []byte) int {
	return bytes.Count(content, []byte("\n"))
}

func CalculateNoOfWords(content []byte) int {
	return len(bytes.Fields(content))
}

func CalculateNoOfChars(content []byte) int {
	return len(bytes.Runes(content))
}