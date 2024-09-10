package main


import ("fmt")

func Loops(){
	// GoLang have only one loop
	// 1. for loop
	// for loop have 3 different types

	// 1. for loop with single condition
	// for condition{
		// 	// code block
		// }

		//  This can be used as a while loop
		i := 0
		for i < 5 {
			// code block
			i++
		}
	


	// 2. for loop with init, condition and post
	// for init; condition; post{
	// 	 code block
	// }
		for j := 0; j < 5; j++ {
			fmt.Println(i)
		}



	// 3. for loop with range
	// for index, value := range array{
	// 	// code block
	// }
	array := []int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}


}


// In Go, the range keyword is used in for loops to iterate over elements in various data structures such as arrays, slices, maps, strings, and channels. 
// It provides a convenient way to loop through the elements and access both the index and the value of each element.