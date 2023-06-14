package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Process started.")
	// Check the current platform
	currentOS := runtime.GOOS
	currentArch := runtime.GOARCH
	fmt.Printf("Running on %s/%s\n", currentOS, currentArch)

	// Print a message after the spinner
	fmt.Println("Process completed.")
	fmt.Println("Welcome to the Multi-Arch World")
}
