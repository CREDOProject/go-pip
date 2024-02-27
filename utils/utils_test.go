package utils

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/CREDOProject/sharedutils/shell"
)

type execShim func() shell.ExecShim

type myLookPath struct {
	LookPathFunc func(string) (string, error)
}

func (m myLookPath) LookPath(name string) (string, error) {
	return m.LookPathFunc(name)
}

func (m myLookPath) Command(cmd string, args ...string) *exec.Cmd {
	return nil
}

func mockExec() execShim {
	return func() shell.ExecShim {
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

func TestLooksLikePip(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"pip3", true},
		{"pip3.9", true},
		{"pip3.10", true},
		{"pip3.10.1", true},
		{"pip3.10.1a", false},    // Invalid, contains characters after the version
		{"pip3.10.1.1", false},   // Invalid, contains multiple dots after the version
		{"pip3.10.1.1.1", false}, // Invalid, contains multiple dots after the version
		{"pip2", false},          // Invalid, should start with "pip3"
		{"pip", false},           // Invalid, should contain version number
		{"pip3exe", false},       // Invalid, should only contain digits and dots after "pip3"
		{"python3", false},       // Invalid, should start with "pip3"
		{"pypip3.10", false},     // Invalid, should start with "pip3"
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			result := looksLikePip(test.filename)
			if result != test.expected {
				t.Errorf("looksLikePip(%s) returned %t, want %t", test.filename, result, test.expected)
			}
		})
	}
}

func TestPipBinaryFrom(t *testing.T) {
	tests := []struct {
		path        string
		mockResult  []string
		mockErr     error
		expected    string
		expectedErr error
	}{
		{
			path:        "/path/to/some/directory",
			mockResult:  nil,
			mockErr:     nil,
			expected:    "",
			expectedErr: errors.New("No pip found."),
		},
		{
			path:        "/path/to/some/directory",
			mockResult:  nil,
			mockErr:     errors.New("Some error occurred"),
			expected:    "",
			expectedErr: errors.New("Some error occurred"),
		},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			// Call PipBinaryFrom with the mock
			result, _ := PipBinaryFrom(test.path)

			// Check if the result matches the expectation
			if result != test.expected {
				t.Errorf("PipBinaryFrom(%s) returned %s, want %s", test.path, result, test.expected)
			}

		})
	}
}
