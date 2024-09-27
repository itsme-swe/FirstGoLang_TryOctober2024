package main

import "fmt"

func main() {
	var username string = "Harsh Mehra"		// declaring variable string type
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n", username);

	var isLoggedIn bool = true	// declaring variable boolean type
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)	// using format specifier to check the data type of variable

	
}