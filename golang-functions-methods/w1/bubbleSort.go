package main

import (
	"fmt"
)

func main() {
	var array [100]int
	numberOfElements := 0
	fmt.Printf("insert the number of integer to be evaluated \n")
	_, _ = fmt.Scanf("%d%*c", &numberOfElements)
	var sli = array[0:numberOfElements]
	buildSliceToBeSorted(numberOfElements, sli)
	bubbleSort(sli)
	fmt.Print(sli)
}

func bubbleSort(sli []int) {
	for i := 1; i < len(sli); i++ {
		for j := 0; j < len(sli)-i; j++ {
			if sli[j] >= sli[j+1] {
				swap(sli, j)
			}
		}
	}

}

func swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

func buildSliceToBeSorted(numberOfElements int, sli []int) {
	for i := 0; i < numberOfElements; i++ {
		element := 0
		fmt.Printf("insert the element %d \n", i+1)
		_, _ = fmt.Scanf("%d%*c", &element)
		sli[i] = element
	}
}
