package main

import "fmt"

var numbers = [int(1e7) + 1]bool{}

func sieve() {
	for j := 2; j < int(1e7) + 1; j++ {
		numbers[j] = true
	}
	for p := 2; p*p < int(1e7) + 1; p++ {
		if numbers[p] {
			for i := p*p; i < int(1e7) + 1; i += p {
				numbers[i] = false
			}
		}
	}
}

func IsPrime(n int) bool {
	return numbers[n]
}

func main() {
	sieve()
	n, x := 0, 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		if IsPrime(x) {
			fmt.Printf("%d eh primo\n", x)
		} else {
			fmt.Printf("%d nao eh primo\n", x)
		}
	}
}
