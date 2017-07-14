package main 

import "fmt"
import "math"

func main() {

	add(75,25)
	sub(75,25)
	mul(75,25)
	div(75,25)
	inc_dec(75,25)	
	abs(-100)
	
}

func add(a , b int) int{
	
	var total int 
	total = a + b
	fmt.Println("\naddtion")
	fmt.Println("the sum a of and b is :",total)
	return total
}

func sub(a , b int) int{
	
	var total int 
	total = a - b
	fmt.Println("\nsubstraction")
	fmt.Println("the difference of a and b is :",total)
	return total
}

func mul(a , b int) int{
	
	var total int 
	total = a * b
	fmt.Println("\nmultiplication")
	fmt.Println("the product of and b is :",total)
	return total
}

func div(a , b float64) float64{
	
	var total float64
	total = a / b
	fmt.Println("\ndivision")
	fmt.Println("the division of and b is :",total)
	return total
}

func inc_dec(a , b int) int{
	
	fmt.Println("\nincremnt and drecrement")
	a++
	b--
	fmt.Println("the increment of a is :",a)
	fmt.Println("the decrement of b is :",b)
	return a 
}

func abs(x float64) float64{	
	fmt.Println("\nthe abs value of",x,"is :",math.Abs(x))
	return x
}