package main

func binarySearch(a []int, n int) bool {
	low := a[0]
	high := a[len(a)-1]

	for high >= low {
		mid := (high + low) / 2
		if a[mid] == n {
			return true
		} else if a[mid] < n {
			low = mid + 1
		} else if a[mid] > n {
			high = mid - 1
		}
	}
	return false
}
