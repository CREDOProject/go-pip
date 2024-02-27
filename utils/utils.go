package utils

import (
	"errors"
	"regexp"

	"github.com/CREDOProject/sharedutils/files"
	"github.com/CREDOProject/sharedutils/shell"
)

const pip = "pip3"

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
	var pipFileRegex = regexp.MustCompile(`^pip3(\.\d\d?)?\.?(\.\d\d?)?$`)
	return pipFileRegex.MatchString(name)
}
