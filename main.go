/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program asks the user to supply their name, age, and whether they are a student (true/false) and tells what life stage they are in.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Declaring variables
	var yourName string
	var ageAsString string
	var age int
	var student string

	// Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")
	yourName, _ = reader.ReadString('\n')
	yourName = strings.TrimSpace(yourName)

	fmt.Print("How old are you? ")
	ageAsString, _ = reader.ReadString('\n')
	ageAsString = strings.TrimSpace(ageAsString)
	age, _ = strconv.Atoi(ageAsString)

	fmt.Print("Are you a student? (true/false) ")
	student, _ = reader.ReadString('\n')
	student = strings.TrimSpace(student)

	// Output
	if student == "true" {
		// Student case
		if age >= 13 && age <= 19 {
			fmt.Printf("Student %s is a teenager.\n", yourName)
		} else if age >= 5 && age <= 12 {
			fmt.Printf("Student %s is a child.\n", yourName)
		} else {
			fmt.Printf("%s is in a different life stage.\n", yourName)
		}
	} else {
		// Not a student case
		if age >= 20 && age <= 30 {
			fmt.Printf("%s is a young adult.\n", yourName)
		} else {
			fmt.Printf("%s is in a different life stage.\n", yourName)
		}
	}
}
