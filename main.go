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
	Name string            `json:"name,omitempty"`
	Env  map[string]string `json:"env,omitempty"`
	// https://github.com/orgs/community/discussions/45342#discussioncomment-4778446
	// The value for uses must be a literal. I'm not aware of any location in the docs that explicitly says so, but
	// I've seen GitHub staff confirm it on this forum, with the rationale that it's needed to enforce restrictions on which actions/workflows can be called.
	Uses string            `json:"uses,omitempty"`
	With map[string]string `json:"with,omitempty"`
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
			{Uses: "gitleaks/gitleaks-action@v2", Name: "golangci-lint", With: map[string]string{"args": "--timeout=3m", "version": "v1.55.2"}},
			{Uses: "golangci/golangci-lint-action@v3", Name: "gitleaks", Env: map[string]string{"GITHUB_TOKEN": "${{ secrets.GITHUB_TOKEN }", "GITLEAKS_LICENSE": "${{ secrets.GITLEAKS_LICENSE}}"}},
		},
	}

	if command == "generate-matrix" {
		bytes, _ := json.Marshal(wfl)
		fmt.Println(string(bytes))
	} else {
		fmt.Println("Invalid command. Usage: go run main.go generate")
	}

}
