package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your mobile number")

	input, _ := reader.ReadString('\n')

	fmt.Println("Your mobile number is: ", input)
	fmt.Printf("The data type of input is %T",input)
}