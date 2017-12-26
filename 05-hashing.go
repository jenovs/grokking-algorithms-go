package main

func simpleHash(s string) int {
	len := 10

	p := primes(52)

	sum := 0
	for _, c := range s {
		if c >= 65 && c <= 90 {
			sum += p[c-65]
		} else if c >= 97 && c <= 122 {
			sum += p[c-97+26]
		}
	}
	return sum % len
}

func primes(n int) []int {
	p := []int{1, 2}

loop:
	for i := 3; len(p) < n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue loop
			}
		}
		p = append(p, i)
	}
	return p
}
