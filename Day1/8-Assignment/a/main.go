package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("file not in the root directory")

func checkError(err error) {
	if err != nil {
		//newErr := fmt.Errorf("%w %w", err, ErrFileNotFound)
		fmt.Println("Error:", err)
		return
	}
}

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

func openFile(fileName string) *os.File{
	f, err := os.Open(fileName)
	checkError(err)
	return f
}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	checkError(err)
	fmt.Println("\n\n--------------Suceess in file deletion-----------------")
}

func main() {
	path := "C:\\Users\\ORR Training 16\\OneDrive\\go-training\\Go-Assignment\\TekSystems\\Day1\\8-Assignment\\a.txt"
	f := openFile(path)

	fileLines := readFile(f)
	for _, line := range fileLines {
		fmt.Println(line)
	}

	removeFile(path)
}
