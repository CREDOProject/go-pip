package utils

import "os/exec"

const pip = "pip3"

var execCommander = newExecShim

func newExecShim() IExecLookPath {
	shim := ExecShim{}
	return shim
}

// Function used to find the pip binary in the system.
func DetectPipBinary() (string, error) {
	return execCommander().LookPath(pip)
}

// Interface abstracting the exec API
type IExecLookPath interface {
	LookPath(string) (string, error)
}

// Shim struct to attach Exec methods
type ExecShim struct{}

// Function used to call the exec API.
func (exc ExecShim) LookPath(name string) (string, error) {
	program, err := exec.LookPath(pip)
	if err != nil {
		return "", err
	}
	return program, nil
}
