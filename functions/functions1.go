package main
import (
	"fmt"
	"errors"
)

func main(){
	fmt.Printf("The Follwing is a simple calculator\n\n")
	var num1 int = 300
	var num2 int = 61
	var Addition int = addition(num1, num2)
	fmt.Printf("The result of adding the integers %v and %v = %v\n", num1, num2, Addition)
	var Subtraction int = subtraction(num1, num2)
	fmt.Printf("The result of subtracting the integers %v and %v = %v\n", num1, num2, Subtraction)
	var Multiplication int = multiplication(num1, num2)
	fmt.Printf("The result of multiplying the integers %v and %v = %v\n", num1, num2, Multiplication)
	var quotient, remainder, err = division(num1,num2)
	if err!=nil{
		fmt.Printf(err.Error())
	} else if remainder==0{
		fmt.Printf("The quotient is %v", quotient)
	} else{
		fmt.Printf("The quotient is %v and the remainder is %v", quotient, remainder)
	}
}

func addition(num1 int, num2 int) int{
	var result int = num1 + num2
	return result 
}

func subtraction(num1 int, num2 int) int {
	var result int = num1- num2
	return result
}

func multiplication(num1 int, num2 int) int{
	var result = num1 * num2
	return result
}

func division(num1 int, num2 int)(int, int, error){
	var err error
	if num2 == 0 {
		err = errors.New("Cannot divide by 0")
		return 0, 0, err
	}
	var result int = num1 / num2
	var remainder int = num1 % num2
	return result, remainder, err
}