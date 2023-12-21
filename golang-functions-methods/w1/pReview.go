package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("Give me 10 integers, separated by space: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	slice := make([]int, 0, 10)
	slice = StringsToInts(strings.Split(input[:len(input)-1], " "), slice)

	BubbleSort(slice)

	fmt.Print("Ta-daaah, sorted! ", slice)
}

func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				Swap(values, j, j+1)
			}
		}
	}
}

func Swap(values []int, i int, j int) {
	temp := values[i]
	values[i] = values[j]
	values[j] = temp
}

func StringsToInts(input []string, slice []int) []int {
	for _, el := range input {
		number, err := strconv.Atoi(el)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		slice = append(slice, number)
	}

	return slice
}
