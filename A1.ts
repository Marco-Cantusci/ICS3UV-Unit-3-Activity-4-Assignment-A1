/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-18
 * @fileoverview Calculate the total cost of nuts, bolts, and washers.
 */

// declare variables
const boltPrice: number = 10;
const nutPrice: number = 5;
const washerPrice: number = 3;

// input

// bolts input
const boltAmountString: string = prompt("How many bolts would you like to purchase? \n") || "0";
const boltAmountNumber: number = parseInt(boltAmountString);


// nuts input
const nutAmountString: string = prompt("How many nuts would you like to purchase? \n") || "0";
const nutAmountNumber: number = parseInt(nutAmountString);

// washer amount
const washerAmountString: string = prompt("How many washers would you like to purchase? \n") || "0";
const washerAmountNumber: number = parseInt(washerAmountString);

// processing

// price calculation
const boltTotalPrice: number = boltAmountNumber * boltPrice;
const nutTotalPrice: number = nutAmountNumber * nutPrice;
const washerTotalPrice: number = washerAmountNumber * washerPrice;
const totalPrice: number = boltTotalPrice + nutTotalPrice + washerTotalPrice;

// print number of each item
console.log(`Number of bolts: ${boltAmountNumber}`);
console.log(`Number of nuts: ${nutAmountNumber}`);
console.log(`Number of washers: ${washerAmountNumber}`);

// if more bolts than nuts
if (boltAmountNumber > nutAmountNumber) {
  console.log("Check the Order - not enough nuts for the bolts you purchased.");

// if more bolts than washers
} else if (boltAmountNumber > washerAmountNumber) {
  console.log("Check the Order - not enough washers for the bolts you purchased.");

// if everything is okay
} else {
  console.log("Your order is OK.");
}

// total price
console.log(`Your total cost of the order is ${totalPrice} cents.`);

console.log("\nDone.");
