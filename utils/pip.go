package utils

import "regexp"

const pip = "pip3"

var pipFileRegex = regexp.MustCompile(`^pip3(\.\d\d?)?\.?(\.\d\d?)?$`)
