package main 

import (
	"fmt"
	"runtime"
	"time"
	)

func main() {
	
	//for loop
	for_loop()

	//for loop -2
	for_loop_2()

	//using for forever
//	forever()

	//using if loop
	if_loop()

	//sqrt finction
	fmt.Println(Sqrt(2))

	//using switch
	switch_func()

	//switch -2
	switch_weekend()

	//defer - defer waits for its surrounding functions to complete first ---then its executes 
	// it acts as LAST-IN-FIRST-OUT
	defer_func()	
}

//typical for loop very similar to c for(i = 0;i<10;i++)
func for_loop() int{
	sum := 0

	fmt.Println("for loop")
	for i:=0; i<10; i++{
		sum += 1
	}
	fmt.Println(sum)
	
	return sum
}


//for loop implimentation -2 
func for_loop_2() int{
	sum := 1
	sum2 := 1

	for ; sum < 1000;{
		sum += sum
	}

	//similar to while loop while(sum<1000)
	for sum2 < 1000{
		sum2 += sum2
	}
	fmt.Println(sum)
	fmt.Println(sum2)
	return sum
}

// using for forever
func forever() int{	
	for{

	}
}

//using if loop
func if_loop() int{
	x := 10
	y := 20

	fmt.Println("\nif loop")
	if x + y < 20{
		fmt.Println("x and y are less than 20")
	}else {
		fmt.Println("x and y are greater than 20")
	}
	return x
}

//sqrt function
func Sqrt(x float64) float64 {
	z := 1.0
	fmt.Println("\nSqrt function")
	for i := 0; i < 100; i++ {
		z = z - (((z*z) - x) / (2 * z))
	}

	return z
}

//using switch
func switch_func() int{	
	
	i := 1
	fmt.Println("\n using switch")
	fmt.Println("go runs on :")	
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s",os)
	}
	return i 
}

//using switch evaluation order or switch -2
func switch_weekend() int{
	
	i := 1
	fmt.Println("\nswitch - 2 ")
	today := time.Now().Weekday()
	fmt.Println("today is ",today)

	switch time.Saturday {
	case today + 0:
		fmt.Println("today is Saturday")
	case today + 1:
		fmt.Println("tomm is Saturday")
	case today + 2:
		fmt.Println("day after is Saturday")
	default:
		fmt.Println("its far away")
	}
	return i
}

//defer function
func defer_func() int{
	
	fmt.Println("\ndefer function")
	i := 1
	defer fmt.Println("world !!!")
	fmt.Println("hello ")
	
	fmt.Println("counting")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	return i
}