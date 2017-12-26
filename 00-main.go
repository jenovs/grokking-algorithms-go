package main

import "fmt"

func main() {
	fmt.Println("01 - Binary search")
	s1 := []int{1, 2, 3, 5, 6, 7}
	result := binarySearch(s1, 4)
	fmt.Println(result)
	result = binarySearch(s1, 3)
	fmt.Println(result)

	fmt.Println("\n02 - Selection sort")
	s2 := []int{9, 2, 5, 2, 8, 4, 1}
	fmt.Println(s2)
	s2 = selectionSort(s2)
	fmt.Println(s2)

	fmt.Println("\n04 - Recursive sum")
	s4 := []int{2, 4, 6, 88, 10, 61, 3, 45}
	fmt.Println(sumRecursive(s4))

	fmt.Println("\n04 - Recursive count")
	fmt.Println(countRecursive(s4))
	fmt.Println(countRecursive([]int{}))

	fmt.Println("\n04 - Recursive find max")
	fmt.Println(findMaxRecursive(s4))
}
