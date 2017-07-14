package main

import "fmt"

func main() {
	
	//pointer
	pointer()

	//structures
	fmt.Println("\nstructures")
	fmt.Println(structures{1,3})

	//usinf strutures in functions
	struct1()
}

//pointers
func pointer() int{
	
	i,j := 91,5584

	fmt.Println("pointer")
	fmt.Println("value of i(before) :",i)
	p := &i                 				  	// point to i
	fmt.Println(*p)     					    // read i through the pointer
	*p = 55            					       	// set i through pointer i
	fmt.Println("value of i(after) :",i)
	
	fmt.Println("value of j(before) :",j)
	p = &j                 						// point to j
	*p = *p / 66								// divide j through pointer p
	fmt.Println("value of j(after)",j)     		// display the new value of j
	return i
}

//structures
type structures struct{
	X int
	Y int
}

//using structures in functions
func struct1() int{
	
	i := 1
	a := structures{55,63}
	fmt.Println("a.X before : ",a.X)
	a.X = 99
	fmt.Println("a.X after : ",a.X)
	fmt.Println("strct value : ",a)
	return i
}