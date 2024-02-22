package utils

import (
	"github.com/CREDOProject/go-pip/shell"
)

const pip = "pip3"

var execCommander = shell.NewExecShim

// Function used to find the pip binary in the system.
func DetectPipBinary() (string, error) {
	return execCommander().LookPath(pip)
}
