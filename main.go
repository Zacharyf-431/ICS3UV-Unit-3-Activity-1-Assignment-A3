/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-12
	* This file displays the multipication table from 0 to 12, with the number 9
	*/

package main

import "fmt"

func main() {
	// assign constant 
	const multiplier float32 = 9

	// INPUT - none 

	// Process
	// assigning the variables of each seperate multiplying number
	var val0 float32 = 0 * multiplier
	var val1 float32 = 1 * multiplier
	var val2 float32 = 2 * multiplier
	var val3 float32 = 3 * multiplier
	var val4 float32 = 4 * multiplier
	var val5 float32 = 5 * multiplier
	var val6 float32 = 6 * multiplier
	var val7 float32 = 7 * multiplier
	var val8 float32 = 8 * multiplier
	var val9 float32 = 9 * multiplier
	var val10 float32 = 10 * multiplier
	var val11 float32 = 11 * multiplier
	var val12 float32 = 12 * multiplier

	// OUTPUT
	fmt.Println("0 X", multiplier, "=", val0)
	fmt.Println("1 X", multiplier, "=", val1)
	fmt.Println("2 X", multiplier, "=", val2)
	fmt.Println("3 X", multiplier, "=", val3)
	fmt.Println("4 X", multiplier, "=", val4)
	fmt.Println("5 X", multiplier, "=", val5)
	fmt.Println("6 X", multiplier, "=", val6)
	fmt.Println("7 X", multiplier, "=", val7)
	fmt.Println("8 X", multiplier, "=", val8)
	fmt.Println("9 X", multiplier, "=", val9)
	fmt.Println("10 X", multiplier, "=", val10)
	fmt.Println("11 X", multiplier, "=", val11)
	fmt.Println("12 X", multiplier, "=", val12)

	fmt.Println("\nDone.")
	}