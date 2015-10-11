package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	fmt.Println("Binary Search")
	arrayzor := []int{0, 2, 4, 5, 6, 7, 8}

	index := binarySearch(arrayzor, 2)

	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("arrayzor[", index, "] = ", arrayzor[index])
	}

}

func binarySearch(arrayzor []int, toFind int) int {
	var low, high int
	low = arrayzor[0]
	high = arrayzor[len(arrayzor)-1]

	if toFind < low || toFind > high {
		return -1
	}

	for low <= high {
		mid := low + (high-low)/2
		switch {
		case toFind < arrayzor[mid]:
			high = mid - 1

		case toFind > arrayzor[mid]:
			low = mid + 1

		case toFind == arrayzor[mid]:
			return mid

		}

	}
	return -1

}
