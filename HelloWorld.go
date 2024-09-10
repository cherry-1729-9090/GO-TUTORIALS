// every go program startes with a package declaration which provides a way for us to reuse code
package main

// fmt (format) is a library that has all the functions for formatting input and output
import (
	"fmt"
)

// every go program has a main function that is the entry point for the executable program
func main() {
	// fmt.Println("Hello World")
	// fmt.Println("Address of a is : ", &a) // This is the variable from Pointers.go

	// Format printing, using the Printf function from format(fmt) library

	var name string = "John Doe"
	const pi float64 = 3.14159265359
	win := true


	// %s is used to format strings
	fmt.Printf("Name: %s\n", name)

	// %f is used to format float numbers
	fmt.Printf("Value of Pi: %f\n", pi)

	// %t is used to format boolean values
	fmt.Printf("Win: %t\n", win)

	// %T is used to format the type of the variable
	fmt.Printf("Type of name: %T\n", name)

	// %v is used to format the value of the variable
	fmt.Printf("Value of pi: %v\n", pi)

	// %c is used to format characters
	fmt.Printf("First letter of name: %c\n", name[0])

	// %d is used to format decimal numbers
	fmt.Printf("Age: %d\n", 23)

	// %b is used to format binary numbers
	fmt.Printf("10 in binary: %b\n", 10)

	// %e is used to format scientific notation
	fmt.Printf("Pi in scientific notation: %e\n", pi)

	// %o is used to format octal numbers
	fmt.Printf("10 in octal: %o\n", 10)

	// %q is used to format single-quoted strings
	fmt.Printf("Name: %q\n", name)

	// %p is used to format the pointer value
	fmt.Printf("Address of name: %p\n", &name)

	// %x is used to format hexadecimal numbers
	fmt.Printf("Address of name in hexadecimal: %x\n", &name)

	// %U is used to format Unicode format
	fmt.Printf("Name in Unicode: %U\n", name)


}

// In every package there is only one main function and one main package


// %v: This is the default format specifier. It prints the value in a default format based on the type of the variable.
// %s: This format specifier is used to print a string. If the variable is not a string, using %s will cause a runtime error.