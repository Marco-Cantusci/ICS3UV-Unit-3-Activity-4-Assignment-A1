/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-18
 * @fileoverview Calculate the total cost of nuts, bolts, and washers.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare variables
	const boltPrice int = 10
	const nutPrice int = 5
	const washerPrice int = 3

	var boltAmountString string
	var boltAmountInt int
	var nutAmountString string
	var nutAmountInt int
	var washerAmountString string
	var washerAmountInt int

	reader := bufio.NewReader(os.Stdin)

	// input

	// bolts input
	fmt.Print("How many bolts would you like to purchase? \n")
	boltAmountString, _ = reader.ReadString('\n')
	boltAmountString = strings.TrimSpace(boltAmountString)

	boltAmountInt, _ = strconv.Atoi(boltAmountString)

	// nut input
	fmt.Print("How many nuts would you like to purchase? \n")
	nutAmountString, _ = reader.ReadString('\n')
	nutAmountString = strings.TrimSpace(nutAmountString)

	nutAmountInt, _ = strconv.Atoi(nutAmountString)

	// washer input
	fmt.Print("How many washers would you like to purchase? \n")
	washerAmountString, _ = reader.ReadString('\n')
	washerAmountString = strings.TrimSpace(washerAmountString)

	washerAmountInt, _ = strconv.Atoi(washerAmountString)

	// processing

	// price calculation
	var boltTotalPrice int = boltAmountInt * boltPrice
	var nutTotalPrice int = nutAmountInt * nutPrice
	var washerTotalPrice int = washerAmountInt * washerPrice
	var totalPrice int = boltTotalPrice + nutTotalPrice + washerTotalPrice

	fmt.Printf("Number of bolts: %d \n", boltAmountInt)
	fmt.Printf("Number of nuts: %d \n", nutAmountInt)
	fmt.Printf("Number of washers: %d \n", washerAmountInt)

	// if more bolts than nuts
	if boltAmountInt > nutAmountInt {
		fmt.Println("Check the Order - not enough nuts for the bolts you purchased.")

	} else if boltAmountInt > washerAmountInt {
		fmt.Println("Check the Order - not enough washers for the bolts you purchased.")

	} else {
		fmt.Println("Your order is OK.")
}

// print total cost
	fmt.Printf("Your total cost of the order is %d cents. \n", totalPrice)

	fmt.Println("\nDone.")
}
