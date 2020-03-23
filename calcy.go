package main

import "fmt"

func main(){
	var ch int
	var (
		a float64
		b float64
	) 
	fmt.Println("I can do all this:")
	fmt.Println("1.Addition")
	fmt.Println("2.Subtraction")
	fmt.Println("3.Multiplication")
	fmt.Println("4.Division")
	fmt.Println("So, what do you want to do today?")
	fmt.Scan(&ch)
	fmt.Println("And on whom do you want to do 'em on? *wink*")
	fmt.Scan(&a)
	fmt.Scan(&b)
	switch(ch){
	case 1:fmt.Println("Result = ", a+b)
	case 2:fmt.Println("Result = ", a-b)
	case 3:fmt.Println("Result = ", a*b)
	case 4:fmt.Println("Result = ", a/b)
	default:fmt.Println("What are you even trying to do?")
	}
}