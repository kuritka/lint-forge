package installer

import (
	"fmt"
	"os/exec"
)

type Runner interface {
	Run(string, ...string) *Output
}

func (i *Output) Run(binaryPath string, args ...string) *Output {
	if i.Error != nil {
		return i
	}
	o := new(Output)
	// Create a command
	cmd := exec.Command(binaryPath, args...)

	// Capture the standard output and standard error
	o.Output, o.Error = cmd.CombinedOutput()
	fmt.Println(o)
	if o.Error != nil {
		o.Error = fmt.Errorf("error: %s", o.Error)
	}
	return o
}
