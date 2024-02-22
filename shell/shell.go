package shell

import "os/exec"

// Interface abstracting the exec API
type IExecLookPath interface {
	LookPath(string) (string, error)
}

// Shim struct to attach Exec methods
type ExecShim struct{}

// Function used to call the exec LookPath API.
func (exc ExecShim) LookPath(name string) (string, error) {
	program, err := exec.LookPath(name)
	if err != nil {
		return "", err
	}
	return program, nil
}

func NewExecShim() IExecLookPath {
	shim := ExecShim{}
	return shim
}
