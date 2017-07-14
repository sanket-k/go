package main 

import (
   "fmt"
   "time"   
   "math/rand"                              // means rand is a subdirectory inside math
   "math"
)

func main() {
	fmt.Println("this is from playground\n")
	fmt.Println("the time is",time.Now())

	fmt.Println("\na random number is = ",rand.Intn(100))        // Intn once set will return the same number 

	fmt.Printf("\nthe squareroot of 2 is : %g", math.Sqrt(2))

	fmt.Println("\nthe value of pi is :",math.Pi)

	//function
	fmt.Println("the sum is : ",add(42,13))

	//swap function
	a,b := swap("hello","world") 
	fmt.Println(a,b)
}	

func add(x , y int) int{
	var total int
	total = x + y
	return total 
}

func swap(x,y string) (string,string){
	return y,x
}