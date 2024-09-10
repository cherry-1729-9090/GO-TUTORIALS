package main

 import ("fmt")


// In Go, an array is a fixed-length sequence of elements of a single type.
func Arrays(){
	// Size of array is always fixed, so we can't change the size of array
	// In here since n may not be a constant, we can't use it as size of array

	// var n int
	// fmt.Scanln(&n)
	// var Numbers [n]int
	// for i := 0;i < n;i++ {
	// 	fmt.Scanln(&Numbers[i])
	// }


	// We can use slice instead of array

	var n int
	fmt.Scan(&n)
	Numbers := make([]int, n)
	for i := 0;i < n;i++ {
		fmt.Scan(&Numbers[i])
	}
	

	EvenNumbers := [5]int {0,2,4,6,8}
	// ways to iterate over arrays 
	// 1. Using for loop
	for i := 0;i < len(EvenNumbers);i++ {
		fmt.Println(EvenNumbers[i])
	}

	// 2. Using range
	// range returns two values, index and value
	for index, value := range EvenNumbers{
		fmt.Println(index, value)
	}


	// Slicing the arrays
	// Slicing is a technique to create a new array by selecting elements from another array
	evenSlice := []int {5,3,4,5,3}
	fmt.Println(evenSlice[1:3]) // [3,4]


	// different methods in slicing
	// 1. make([]T, length, capacity)
	// It creates a slice of type T with a length and capacity specified by length and capacity arguments.
	//  In Go, when you create a slice with a specified capacity using the make function, 
	// the slice has an initial length and an allocated capacity.
	// The capacity determines how many elements the slice can hold before it needs to allocate new memory to accommodate additional elements.
	arr1 := make([]int, 5, 10)
	fmt.Println(arr1) // [0 0 0 0 0]

	// 2. append(slice, elements)
	// It appends elements to the end of a slice and returns the resulting slice.
	arr2 := []int {1,2,3,4,5}
	arr2 = append(arr2, 6)
	fmt.Println(arr2) // [1 2 3 4 5 6]

	// 3. copy(destination, source)
	// It copies elements from a source slice to a destination slice and returns the number of elements copied.
	arr3 := []int {1,2,3,4,5}
	arr4 := make([]int, 5)
	copy(arr4, arr3)

	// 4. len(slice)
	// It returns the length of the slice.
	fmt.Println(len(arr4)) // 5

	// 5. cap(slice)
	// It returns the capacity of the slice.
	fmt.Println(cap(arr4)) // 5

	// 6. slice[low:high]
	// It returns a slice of the slice from index low to high-1.
	fmt.Println(arr4[1:3]) // [2 3]

	


}