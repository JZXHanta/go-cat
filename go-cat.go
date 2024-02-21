package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFile(filepath string) (arrayOfStrings []string) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fileString := string(file)
	arrayOfStrings = strings.Split(fileString, "\n")
	return
}

// Renders input text with line numbers on non-blank lines
func renderB(lineArray []string) {
	num := 1
	for _, line := range lineArray {
		if strings.TrimSpace(line) == "" {
			fmt.Printf("%s\n", line)
		} else {

			fmt.Printf("%d %s\n", num, line)
			num += 1
		}
	}
}

// Renders input text with line numbers on all lines, even blank
func renderN(lineArray []string) {
	for num, line := range lineArray {
		fmt.Printf("%d %s\n", num, line)
	}
}

// Renders input text with no additions
func render(lineArray []string) {
	str := strings.Join(lineArray, "\n")
	// for _, line := range lineArray {
	// 	fmt.Printf("%s\n", line)
	// }
	fmt.Print(str)
}

func readPipe() (arrayOfStrings []string) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	str := string(stdin)
	arrayOfStrings = strings.Split(str, "\n")
	return
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		log.Fatal("Not enough arguments: expected 1 argument, recieved 0.")
	}

	var filepaths []string
	var str []string

	for _, arg := range arguments {
		switch arg {
		case "-n", "--number", "-b", "--number-nonblank", "-h", "--help":
			continue
		default:
			filepaths = append(filepaths, arg)
		}
	}

	for _, f := range filepaths {
		if f == "-" {
			str = append(str, readPipe()...)
		} else {
			s := readFile(f)
			str = append(str, s...)

		}
	}

	switch arguments[0] {
	case "-n":
		renderN(str)
		os.Exit(0)
	case "-b":
		renderB(str)
	case "--help", "-h":
		fmt.Print(manPage)
		os.Exit(0)
	default:
		render(str)
		os.Exit(0)
	}
}
