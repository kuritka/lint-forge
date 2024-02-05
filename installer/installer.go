package installer

import (
	"errors"
	"fmt"
	"os/exec"
)

type BrewInstaller struct{}

func NewBrewInstaller() *BrewInstaller {
	return &BrewInstaller{}
}

func (b *BrewInstaller) Install(what string) *Output {
	var o = new(Output)
	if what == "" {
		return &Output{[]byte{}, errors.New("no package to install specified")}
	}
	binaryPath := "brew"
	args := []string{"install", what}

	// Create a command
	cmd := exec.Command(binaryPath, args...)

	// Capture the standard output and standard error
	o.Output, o.Error = cmd.CombinedOutput()
	fmt.Println(o)
	if o.Error != nil {
		o.Error = fmt.Errorf("brew error: %s", o.Error)
	}
	return o
}
