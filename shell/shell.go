package shell

import "os/exec"

// Interface abstracting the exec API
type IExecLookPath interface {
	LookPath(string) (string, error)
	Command(string, ...string) *exec.Cmd
}

// Shim struct to attach Exec methods
type ExecShim struct {
}

// Function used to call the exec LookPath API.
func (exc ExecShim) LookPath(name string) (string, error) {
	return exec.LookPath(name)
}

func (exc ExecShim) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

func NewExecShim() IExecLookPath {
	return ExecShim{}
}
