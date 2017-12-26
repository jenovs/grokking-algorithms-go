package main

func selectionSort(s []int) []int {
	sorted := []int{}

	for len(s) > 0 {
		ind := findMin(s)
		sorted = append(sorted, s[ind])
		if len(s) == 1 {
			return sorted
		}
		s = append(s[:ind], s[ind+1:]...)
	}
	return sorted
}

func findMin(s []int) int {
	if len(s) < 1 {
		return -1
	}
	min := s[0]
	ind := 0
	for i := range s {
		if s[i] < min {
			min = s[i]
			ind = i
		}
	}
	return ind
}
