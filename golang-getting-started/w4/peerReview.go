package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	fmt.Println("Enter file name:")
	var fileName string
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) >= 2 {
			names = append(names, Name{fname: parts[0], lname: parts[1]})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for _, name := range names {
		fmt.Println("First Name:", name.fname, ", Last Name:", name.lname)
	}

}
