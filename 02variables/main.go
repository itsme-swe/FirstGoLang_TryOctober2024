package main

import "fmt"

var ExportVariable string = "I am visible outside my package."		// Declaring variable starting with capital letter means its an "public variable".

func main() {
	var username string = "Harsh Mehra"		// declaring variable string type
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n", username);	// using format specifier to check the data type of variable

	var isLoggedIn bool = true	// declaring variable boolean type
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)	

	// default values and some aliases
	var anotherVariable int 
	fmt.Println("Empty value of int ",anotherVariable)	
	/*
	output: 0 
	whenever we declare any variable without value it will return 0.
	*/

	var newName string
	fmt.Println(newName)
	
}