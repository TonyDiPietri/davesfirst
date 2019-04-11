package main

import "fmt"

func main() {
	var thisNumber float32
	var thatnumber float32
	
	thisNumber=100.75
	thatnumber=2.7
	
	fmt.Printf("%-20s%.2f\n","That number is: ", thatnumber)
	fmt.Printf("%-20s%.2f\n","This number is: ", thisNumber)
	fmt.Printf("%-20s%.2f\n","Addition: ", thisNumber+thatnumber)
	fmt.Printf("%-20s%.2f\n","Multiplication: ",thisNumber*thatnumber)
	fmt.Printf("%-20s%.2f\n","Minus: ",thisNumber-thatnumber)
	fmt.Printf("%-20s%.2f\n","Division: ",thisNumber/thatnumber)
	//fmt.Printf("%-20s%.0f\n","modi: ",thisNumber%thatnumber)

	}
