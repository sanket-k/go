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

	//arrays
	array()

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

func array() int{
	
	i := 1
	fmt.Println("\nArrays")
	//array format ==> [n]t ===> n = number of values; t = type 
	//if array size is less than specified then the empty part is filled with zeros
	array_a := [5]int{2,6,3,4}
	fmt.Println(array_a)

	array_str := [2]string{"hello","world !!!"} 
	fmt.Println(array_str)
	
	//array literals ==> array without length
	l := []int {55,2,4,8,6,4}
	fmt.Println(l)

	fmt.Println("length of l = ",len(l))
	fmt.Println("capacity of l = ",cap(l))

	//nil slice
	empty_array := []int{}
	fmt.Println("empty_array len = ",len(empty_array))
	fmt.Println("empty_array cap = ",cap(empty_array))
	if empty_array == nil {
		fmt.Println("the empty_array is empty")
	}

	//using make
	a := make([]int,5)   	// len(a) = 5
	b := make([]int,0,5)	// len(b) = 0 ; cap(b) = 5
	fmt.Println(a)
	fmt.Println(b)
	//printSlice("a",a)
	//printSlice("b",b)

	//make new array
	//array_new := make([],5,5) 
	return i
}