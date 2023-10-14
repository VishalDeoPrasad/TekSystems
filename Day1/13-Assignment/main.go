package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(file *os.File) []string {
	fileScanner := bufio.NewScanner(file) //creating a scanner file
	fileScanner.Split(bufio.ScanLines)    //splitting the file into lines

	var fileLines []string
	for fileScanner.Scan() { //loop through the line one by one
		fileLines = append(fileLines, fileScanner.Text())
	}
	file.Close()
	return fileLines
}

func openFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		error := errors.New("file not found")
		log.Println(path, ",", error)
		os.Exit(1)
	}
	return f
}
func CountWords(contents []string) (int, error) {
	err := errors.New("your file contain nothing or just whitespace")
	if len(contents) == 0 || strings.TrimSpace(contents[0]) == "" {
		fmt.Println("Your .txt file is empty:")
		return 0, err
	}
	cnt := 0
	for _, line := range contents {
		for _, char := range line {
			if char == 32 {
				cnt++
			}
		}
	}
	return cnt, nil
}

func main() {
	path := `sample1.txt`
	file := openFile(path)
	contents := readFile(file)
	cnt, err := CountWords(contents)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Number of words:", cnt)
}
