package main

// As we know every variable has a memory address, we can get the address of the variable by using " & (ampersand) " operator


// Pointers are mainly used to reduce the memory usage and increase the performance of the program
// We pass pointers inside the functions instead of passing the actual values
// This way, we can avoid the creation of a copy of the variable in the memory


import ("fmt")
var a int = 10
func execut(){
	a = 10
	fmt.Println("Value of a is: ", a)
	fmt.Println("Address of a is: ", &a)

	// If we send a variable to a function, it will create a copy of variable in the different memory location,
	// so the original variable will remain unchanged
	changeValue(a)
	fmt.Println("Value of a is: ", a) // 10
	// In the above line, the value of a will be 10,
	// because the changeValue function will create a copy of a in the different memory location


	// Now if we send the address of the variable to a function, it will change the value of the original variable
	changeValueByPointer(&a) 
	fmt.Println("Value of a is: ", a) // 20

	// To send the address of a variable to a function, we need to use the " * (asterisk) " operator
	// This operator is used to get the value of the address
	// To receive the address of a variable in a function, we need to use the " * (asterisk) " operator
	// This operator is used to get the address of the variable


}

func changeValue(x int){
	x = 20
}

func changeValueByPointer(x* int){
	*x = 20
}


// Pointers in Go have data types. 
// A pointer's type is determined by the type of the variable it points to. 
// The type of a pointer is denoted by an asterisk (*) followed by the type of the variable it points to.