package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	var result = make([]Person, 0, 10000)
	var fileName string
	fmt.Printf("Insert the File Name: \n")
	_, err := fmt.Scanf("%s\n", &fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fullname := scanner.Text()
		fnlen := len(fullname)
		spaceLocation := strings.Index(fullname, " ")
		if spaceLocation != -1 {
			result = append(result, Person{FirstName: fullname[0:spaceLocation], LastName: fullname[spaceLocation+1 : fnlen]})
		}

	}
	for _, p := range result {
		fmt.Println(p)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
