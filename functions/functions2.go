package main
import (
	"fmt"
	"math"
)

func Prime(){
	fmt.Printf("Check if a number is prime or not")
	// Prompt the user to enter a number
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Test the isPrime function with the user-input number
	if isPrime(num) {
		fmt.Printf("%v is a prime number.\n", num)
	} else {
		fmt.Printf("%v is not a prime number.\n", num)
	}

}
func isPrime(num int) bool{
	//A prime numbe is a natural number that is greater than 1 and has no positive divisors except from one and itself
	if num <=1 {
		return false
	}
	//Iterate from 2 to the square root of n
	for i := 2 ; i <= (int(math.Sqrt(float64(num)))) ; i++ {
		// If n is divisible by i, it's not prime
		if num % i == 0 {
			return false
		}
	}
	return true
}
//Explanation
// If n is not a prime number, it can be expressed as the product of two numbers, a and b, where both a and b are greater than 1.

// If both a and b are greater than the square root of n, their product would be greater than n. Therefore, at least one of the factors (a or b) must be less than or equal to the square root of n.

// By iterating up to the square root of n in the loop, we efficiently check for divisibility by potential factors without redundant checks beyond the square root. This reduces the number of iterations, making the primality check more efficient.

// In mathematical terms, if n is not a prime number, there exists at least one pair of integers (a, b) such that 1 < a <= b and a * b = n. At least one of a or b must be less than or equal to the square root of n. Therefore, checking up to the square root is sufficient to determine whether n is prime or not.