package main

import "fmt"

func main() {
	fmt.Println("01 - Binary search")
	s1 := []int{1, 2, 3, 5, 6, 7}
	fmt.Printf("%v in %v: %v\n", 4, s1, binarySearch(s1, 4))
	fmt.Printf("%v in %v: %v\n", 3, s1, binarySearch(s1, 3))

	fmt.Println("\n02 - Selection sort")
	s2 := []int{9, 2, 5, 2, 8, 4, 1}
	fmt.Printf("Before sort: %v\n", s2)
	fmt.Printf("After sort: %v\n", selectionSort(s2))

	fmt.Println("\n04 - Recursive sum")
	s4 := []int{2, 4, 6, 88, 10, 61, 3, 45}
	fmt.Printf("Sum of all numbers in %v: %v\n", s4, sumRecursive(s4))

	fmt.Println("\n04 - Recursive count")
	fmt.Printf("Number of elements in %v: %v\n", s4, countRecursive(s4))
	fmt.Printf("Number of elements in %v: %v\n", []int{}, countRecursive([]int{}))

	fmt.Println("\n04 - Recursive find max")
	fmt.Printf("Max value in %v: %v\n", s4, findMaxRecursive(s4))

	fmt.Println("\n04 - Recursive binary search")
	fmt.Printf("%v in %v: %v\n", 4, s1, binarySearchRecursive(s1, 4))
	fmt.Printf("%v in %v: %v\n", 3, s1, binarySearchRecursive(s1, 3))

	fmt.Println("\n05 - Hash tables")
	fmt.Printf("Hash of Esther: %v\n", simpleHash("Esther"))
	fmt.Printf("Hash of Ben: %v\n", simpleHash("Ben"))
	fmt.Printf("Hash of Bob: %v\n", simpleHash("Bob"))
	fmt.Printf("Hash of Dan: %v\n", simpleHash("Dan"))
	fmt.Printf("Hash of Maus: %v\n", simpleHash("Maus"))
	fmt.Printf("Hash of Fun Home: %v\n", simpleHash("Fun Home"))
	fmt.Printf("Hash of Watchmen: %v\n", simpleHash("Watchmen"))
}
