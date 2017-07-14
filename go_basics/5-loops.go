package main 

import "fmt"

func main() {
	
	//for loop
	for_loop()

	//for loop -2
	for_loop_2()

	//using for forever
//	forever()
}

//typical for loop very similar to c for(i = 0;i<10;i++)
func for_loop() int{
	sum := 0
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

