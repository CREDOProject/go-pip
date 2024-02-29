package utils

import (
	"errors"

	"github.com/CREDOProject/sharedutils/files"
	"github.com/CREDOProject/sharedutils/shell"
)

var execCommander = shell.New

// Function used to find the pip binary in the system.
func DetectPipBinary() (string, error) {
	return execCommander().LookPath(pip)
}

func PipBinaryFrom(path string) (string, error) {
	execs, err := files.ExecsInPath(path, looksLikePip)
	if err != nil {
		return "", err
	}
	if len(execs) < 1 {
		return "", errors.New("No pip found.")
	}

	return execs[0], err

}

// looksLikePip returns true if the given filename looks like a Python
// executable.
func looksLikePip(name string) bool {
	return pipFileRegex.MatchString(name)
}
