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
