// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func swap(slice *[]int, i int) {
	temp := (*slice)[i]
	(*slice)[i] = (*slice)[i+1]
	(*slice)[i+1] = temp
}

func bubbleSort(slice *[]int) {
	sort.Ints(*slice)

	length := len(*slice)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if (*slice)[j] > (*slice)[j+1] {
				// swap
				swap(slice, j)
			}
		}
	}
}

func main() {
	var numSlice []int

	fmt.Println("Type up to 10 integers separated by a comma. Integers after 10th number will be ignored.")

	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')

	for i, v := range strings.Split(inputStr, ",") {
		if i < 10 {
			// trim
			trimmedCharacter := strings.TrimSpace(v)
			// to number
			num, _ := strconv.Atoi(trimmedCharacter)
			numSlice = append(numSlice, num)
		}
	}

	fmt.Println("Before sort: ", numSlice)

	bubbleSort(&numSlice)

	fmt.Println("After sort: ", numSlice)
}
