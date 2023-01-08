package main

import "fmt"

// Responsible for the instantiation and startup of the app
func Run() error {
	fmt.Println("Startup...")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
