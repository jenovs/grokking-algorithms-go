package main

func sumRecursive(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[0] + sumRecursive(s[1:])
}

func countRecursive(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return 1 + countRecursive(s[1:])
}

func findMaxRecursive(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return findMax(s, s[0])
}

func findMax(s []int, m int) int {
	if len(s) == 0 {
		return m
	}
	max := m
	if s[0] > m {
		max = s[0]
	}

	return findMax(s[1:], max)
}

func binarySearchRecursive(s []int, n int) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return s[0] == n
	}

	m := (len(s) - 1) / 2
	if s[m] == n {
		return true
	} else if s[m] > n {
		return binarySearchRecursive(s[:m], n)
	} else {
		return binarySearchRecursive(s[m:], n)
	}
}
