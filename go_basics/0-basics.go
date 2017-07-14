package main 

import (
   "fmt"
   "time"   
   "math/rand"                              // means rand is a subdirectory inside math
   "math"
)
//numeric constants
const (	
	// create a huge number by shifting 1 bit left 100 places
	// binary number 1 followed by 100 zeros
 	Big = 1 << 100  

 	//shift right again 100 
 	small = Big >> 99
)

//variable declaration can be at package or function level
var boo bool

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

	//return without arguments
	fmt.Println(spilt(17))

	// variable at package or function level
	var boo1 int
	fmt.Println("\nvariable at package or function level")
	fmt.Println("boo1 is int :",boo1,"\nboo is bool :",boo)

	//short variable declaration
	short(12,54)

	//type conversions
	type_conv()

	//type_inference
	type_inf()

	//numeric constants
	fmt.Println("\nNumeric constants")
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(Big))

}	

func add(x , y int) int{
	var total int
	total = x + y
	return total 
}

func swap(x,y string) (string,string){
	return y,x
}

func spilt(sum int) (x , y int){
	x = sum * 4 / 9
	y = sum - x

	return
}

//short variable
func short(a,b int) (int,int){
	var x,y int = a,b
	z := x+y          //only WORKS in function                               //:= short assignment statement can be used in place of a var declaration with implicit type
	fmt.Println("val of a = ",x,"\nval of b = ",y)
	fmt.Println("val of z and sum of x and y is = ",z)
	return x,y
}

//type conversions 
func type_conv() int{
	var i int = 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i,f,u)
	return i	
}

//type inferrence
func type_inf() int{		
	i := 42         	// it automatically considers it as int 
	f := 3.421   		// it automatically considers it as float64
	g := 0.658 + 0.5i 	// it is considered as complex128

	fmt.Println("\nthis is type inference")
	fmt.Printf("i is of type %T\n",i)
	fmt.Println("i =",i)
	fmt.Printf("f is of type %T\n",f)
	fmt.Println("f =",f)
	fmt.Printf("g is of type %T\n",g)
	fmt.Println("g =",g)

	return i
}

//numeric constants 
func needInt(x int) int{
	return x*10 + 1
}
func needFloat(x float64) float64{
	return x * 0.1
}

