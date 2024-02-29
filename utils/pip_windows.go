//go:build windows

package utils

const pip = "pip3.exe"

var pipFileRegex = regexp.MustCompile(`^pip3(\.\d\d?)?\.?(\.\d\d?)?\.exe$`)
