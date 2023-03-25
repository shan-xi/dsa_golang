package main

import "fmt"

func modifySlice(s *[]int) {
	// Declare a new slice
	newSlice := []int{1, 2, 3}

	// Assign the new slice to the input slice using "&" operator
	s = &newSlice
	(*s)[0] = 100
}

func main() {
	// Declare a slice
	slice := []int{10, 20, 30}

	// Print the original slice
	fmt.Println("Original slice:", slice)

	// Call the modifySlice function and pass the slice as an argument
	modifySlice(&slice)

	// Print the updated slice
	fmt.Println("Updated slice:", slice)
}
