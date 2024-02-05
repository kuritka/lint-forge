// package main
package main

import (
	"encoding/json"
	"fmt"
	"lint-forge/installer"
	"os"
)

type Workflow struct {
	LintJob []Step `json:"include"`
}

type Step struct {
	Project string            `json:"project"`
	Config  string            `json:"config"`
	Args    []string          `json:"args"`
	Name    string            `json:"name,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
}

func main() {

	if len(os.Args) < 2 {
		o := installer.NewBrewInstaller().Install("golangci-lint").Run("golangci-lint", "run")
		if o.Error != nil {
			fmt.Println("Error:", o.Error)
		}
		return
	}
	command := os.Args[1]

	wfl := Workflow{
		LintJob: []Step{
			{Project: "foo", Config: "Debug"},
			{Project: "bar", Config: "Release"},
		},
	}

	if command == "generate-matrix" {
		bytes, _ := json.Marshal(wfl)
		fmt.Println(string(bytes))
	} else {
		fmt.Println("Invalid command. Usage: go run main.go generate")
	}

}
