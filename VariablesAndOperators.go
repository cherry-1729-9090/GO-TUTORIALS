package main

// GoLang have 4 different data types
// 1. Numeric - int, float, complex
// 2. string - string
// 3. boolean - bool
// 4. Derived - array, slice, map, struct, pointer, function, interface, channel

// GoLang have 3 different variable declaration
// 1. var - var name type = value
// 2. shorthand - name := value ( only works inside a function )
// 3. const - const name type = value

// GoLang have 3 different variable scope
// 1. package level - declared outside of any function
// 2. function level - declared inside of a function
// 3. block level - declared inside of a block

// GoLang have 3 different variable visibility
// 1. exported - variable name starts with a capital letter
// 2. unexported - variable name starts with a small letter
// 3. block level - variable name starts with a small letter



// GoLang have same operators as other programming languages

// In goLang a package is just like a project in other programming languages
// In here we can share a variable or a function between different files in the same package
// To share a variable in the package level to any file, we need to declare this variable and initialize it outside of any functions in any file in package

func varAndOperator(){
	// var declaration
	var a int = 10
	var b int = 20
	var c int = 0

	// shorthand declaration
	// := automatically assigns the type of the value
	d := 30
	e := 40
	f := 0

	// const declaration
	var g int = 50
	const h int = 60
	var i int = 0

	// arithmetic operators
	c = a + b
	f = d - e
	i = g * h

	// bitwise operators
	c = a & b
	f = d | e
	i = g ^ h

	// assignment operators
	c += a
	f -= d
	i *= g

	// other operators
	a++
	d--
	g = ^h

	// const declaration
	const j int = 70

}


// In Go, declaring a variable and not using it will cause a compilation error. 
// Go enforces this rule to help keep the code clean and to avoid potential bugs that might arise from unused variables.