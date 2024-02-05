// package main
package main

import (
	"fmt"
	"lint-forge/installer"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		o := installer.NewBrewInstaller().Install("golangci-lint").Run("golangci-lint", "run")
		if o.Error != nil {
			fmt.Println("Error:", o.Error)
		}
		return
	}
	command := os.Args[1]
	if command == "generate-matrix" {
		values := []int{1, 2, 3, 4, 5}
		fmt.Printf("::set-output name=matrix::%v", values)
	} else {
		fmt.Println("Invalid command. Usage: go run main.go generate")
	}

}
