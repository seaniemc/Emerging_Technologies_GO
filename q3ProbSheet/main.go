// q3ProbSheet project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')

	n := 0

	/*Rune is a Type. It occupies 32bit and is meant to represent a Unicode CodePoint.
	  As an analogy the english characters set encoded in 'ASCII' has 128 code points.
	  Thus is able to fit inside a byte (8bit). From this (erroneous) assumption C treated
	  characters as 'bytes' char, and 'strings' as a 'sequence of characters' char*.
	  But guess what. There are many other symbols invented by humans other than the 'abcde..'
	  symbols. And there are so many that we need 32 bit to encode them.
	  In golang then a string is a sequence of runes.
	*/
	rune := make([]rune, len(input))

	//A range clause provides a way to iterate over an array, slice, string, map, or channel.
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	fmt.Println(output)
}

//http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
//http://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
//http://stackoverflow.com/questions/19310700/what-is-a-rune
