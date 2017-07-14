package main                                                   //import main

// import fmt package for display
import "fmt"

func main() {                                                   // open main function 
	var x int                                                   // assign x to int as a variable                         
	var y int                                                   // assign y to int as a variable
	var x1 float64                                              // assign x1 to float as a variable
	var z int                                                   // assign z to int as a variable
	var ham = "hello ham"                                       // assign ham a value of string as a variable

	x = 55                                                      // assign x to 55
	y = 54                                                      // assign y to 54
	x1 = 2.56416584165416514                                    // assign x1 to 2.5...
	z = 65                                                      // assign 65 

	fmt.Println("\nvalue of z is :",z)                          // print value of z 
	fmt.Println(x1)                                             // print value of x1
	fmt.Println("value of x is :",x,"\nvalue of y is :",y)      // print value x and y
	fmt.Println(ham)                                            // print ham

	constant()                                                	// call function constant
	not_const()													// call functoin not-const to display var
}

func constant() {                                               // create a function the use of const     
 	 
	const n = 5555                                              // here n is assigned to 5555 

	//n = 100       //this gives an error                       // n is assigned a different variable   
	fmt.Println("\nthis is a function for const")
	fmt.Println("the value of n is :",n)                        // display the vaule of n
}

func not_const() { 												// create  afunction to display the use of var over const	
	
	var x int = 5555											// asssign x to 5555 as variable		

	x = 100														// reassign x to 100

	fmt.Println("\nthis is a function for not-const")
	fmt.Println("the value of x is :",x)         				// disply the value of x
}
