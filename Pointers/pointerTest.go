package main

import ("fmt")

func main(){
	var a int = 10
	var b* int = &a
	x := *b
	fmt.Println("Value of a is: ", a)
	fmt.Println("Address of a is: ", &a)
	fmt.Println("Value of b is: ", b)
	x = 50
	fmt.Println("Value of x is: ", x)
	fmt.Println("Value of a after changing : ", a)	

}