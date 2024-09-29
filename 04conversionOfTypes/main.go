package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our app")

	fmt.Println("Please rate our food from 1 to 10")
	
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thankyou for your rating, ", input)

	// conversion string to integer and adding 1 to it and we are doing this bcoz we cannot add value to int.
	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Here we are handling error the conditions means that if something is there inside error so throw an error.
	if error != nil {
		fmt.Println(error)
	} else{
		fmt.Println("Added 1 to the rating: ", numRating + 1)
	}

	fmt.Printf("The type variable is: %T", numRating)
	
}