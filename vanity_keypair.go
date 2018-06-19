package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stellar/go/keypair" // https://medium.com/@alexanderleon/stellar-tutorial-making-a-payment-with-go-on-the-testnet-4d0ab2ba0887
)

func main() {
	// While true, calculate a random address, check it for the target pattern, and print the keypair if it matches.
	count := 1
	target := os.Args[1]                // https://gobyexample.com/command-line-arguments
	targetLength := len([]rune(target)) // https://stackoverflow.com/questions/12668681/how-to-get-the-number-of-characters-in-a-string
	printPrettyStuff("Trying to find your address containing: %s\r\n", target)
	// Begin 'while' loop
	for 1 == 1 {
		thisAddress, err := keypair.Random()
		if err != nil {
			log.Fatal(err) // https://gobyexample.com/if-else
		}
		if checkAddress(thisAddress.Address(), target, targetLength) {
			fmt.Printf("Here's the key, keep it secret: \r\n%s\r\n\r\n", thisAddress.Seed())
			break // http://golangtutorials.blogspot.com/2011/06/control-structures-go-for-loop-break.html
		}
		count++
		// Print the number of attempts every 1000 times.
		if count%1000 == 0 {
			fmt.Println(count)
		}
	}
	fmt.Println("Enjoy your personalised address!")
}

func checkAddress(address string, target string, targetLength int) bool {
	// Checks the entire length of an address for an occurrence of the desired pattern.
	// address: the account
	// target: the pattern to match
	// targetLength: pre-calculated so as to not perform the same calculation unnecessarily.
	// returns: bool
	pos := 3
	for pos < (56 - targetLength) {
		chk := address[pos:(pos + targetLength)]
		if chk == target {
			printPrettyStuff("Winner, winnner, chicken dinner! \r\n%s\r\n", address)
			return true
		}
		pos++
	}
	return false
}

func printPrettyStuff(message string, input string) {
	// A simple function for printing information in a nice way.
	// message: The message to be passed to fmt.Printf, with substitutions.
	// input: maximum one substitution.
	border := "============================================================"
	fmt.Println(border)
	fmt.Printf(message, input)
	fmt.Println(border)
}
