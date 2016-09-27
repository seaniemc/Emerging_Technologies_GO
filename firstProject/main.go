// firstProject project main.go
package main

import (
	"fmt"
)

func main() {
	var total int = 0

	for num := 1; num <= 1000; num++ {
		if num%3 == 0 {
			total += num
		}
		if num%5 == 0 {
			total += num
		}
	}
	fmt.Println("The Total number ", total)
}
