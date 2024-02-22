package utils

import (
	"testing"

	"github.com/CREDOProject/go-pip/shell"
)

type execShim func() shell.IExecLookPath

type myLookPath struct {
	LookPathFunc func(string) (string, error)
}

func (m myLookPath) LookPath(name string) (string, error) {
	return m.LookPathFunc(name)
}

func mockExec() execShim {
	return func() shell.IExecLookPath {
		shim := myLookPath{
			LookPathFunc: func(name string) (string, error) {
				return name, nil
			},
		}
		return shim
	}
}

func Test_detectPipBinary(t *testing.T) {
	prevCommander := execCommander
	defer func() { execCommander = prevCommander }()

	execCommander = mockExec()

	result, error := DetectPipBinary()

	if error != nil {
		t.Fatalf("Logic error.")
	}

	if result != pip {
		t.Fatalf("Unexpected result. Got %s, wants %s", result, pip)
	}
}
