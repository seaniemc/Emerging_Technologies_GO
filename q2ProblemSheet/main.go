// q2ProblemSheet project main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	var targetNum int = 10001
	var num = 3.0
	var status int = 1
	var i int = 2
	var j int64

	for i; i <= targetNum; {
		for j := 2; j <= int64(math.Sqrt(num)); j++ {
			if num%j == 0 {
				status = 0
				break
			}
			if status != 0 {
				if i == targetNum {
					fmt.Println(num, "Is the", targetNum, "Prime Number")
				}
				i++
			}
			status = 1
			num++
		}
	}
}
