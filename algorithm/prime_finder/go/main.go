package main

import "fmt"

var MAX = int64(100000000)

func main() {
	primes := GenerateBig(MAX)
	limit := len(primes)

	// fmt.Println(primes[limit-1])

	row := int64(25)
	for i := int64(0); i < int64(limit); i += row {
		for j := int64(0); j < row; j++ {
			index := i + j
			if index >= int64(limit) {
				break
			}
			fmt.Printf("%8d ", primes[index])
		}
		fmt.Print("\n")
	}

}
