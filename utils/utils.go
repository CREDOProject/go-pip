package utils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/CREDOProject/go-pip/shell"
)

const pip = "pip3"

var execCommander = shell.NewExecShim

// Function used to find the pip binary in the system.
func DetectPipBinary() (string, error) {
	return execCommander().LookPath(pip)
}

func PipBinaryFrom(path string) (string, error) {
	execs, err := execsInPath(path)
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
	var pipFileRegex = regexp.MustCompile(`^pip3(\d(\.\d\d?)?)?$`)
	return pipFileRegex.MatchString(name)
}

// execsInPath returns a list of executables in the given path.
// The returned paths are absolute.
//
// The given path should be an absolute path to a directory. If it's not
// a directory, the function will not proceed and return a nil slice.
func execsInPath(path string) ([]string, error) {
	if !isDir(path) {
		return nil, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var execs []string

	for _, entry := range entries {
		if entry.IsDir() || !looksLikePip(entry.Name()) {
			continue
		}

		resolvedPath, err := filepath.EvalSymlinks(filepath.Join(path, entry.Name()))
		if err != nil {
			return nil, err
		}
		info, err := os.Stat(resolvedPath)
		if err != nil {
			return nil, err
		}
		if !isExecutable(info) {
			continue
		}
		execs = append(execs, resolvedPath)
	}

	return execs, nil
}

// isExecutable returns true if the given file info is executable.
// On Windows, it just checks if the file extension is ".exe" or not.
func isExecutable(info fs.FileInfo) bool {
	if runtime.GOOS == "windows" {
		return strings.ToLower(filepath.Ext(info.Name())) == ".exe"
	}
	return info.Mode().IsRegular() && info.Mode()&0o111 != 0
}

// isDir returns true if the given path exists and is a directory.
// It delegates to os.Stat and FileInfo.isDir.
func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
