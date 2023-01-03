package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchAndReplace(filename, search, replace string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), search, replace, -1) // -1 will check here for all the occurences till the very end of the slice
		lines = append(lines, line)
	}
	return ReplaceText(filename, lines)
}

func ReplaceText(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {


	filename := "text.txt" //relative position of the file
	search := "of" // word to search
	replace := "fo" //word to replace with

	SearchAndReplace(filename, search, replace)

	defer fmt.Println(filename,"has been updated, ",search,"has been changed to",replace)
}
