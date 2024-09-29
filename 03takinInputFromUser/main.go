package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcomeMsg string = "Welcome to user input panel:"
	fmt.Println(welcomeMsg)

	// Here we are using bufio and os package to read something 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comma ok || comma error Syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("The data type of rating is %T", input)


}