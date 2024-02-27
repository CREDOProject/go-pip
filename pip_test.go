package gopip

import (
	"strings"
	"testing"
)

func Test_New(t *testing.T) {
	_, err := New("pip3").Seal()
	if err != ErrNoVerb {
		t.Errorf("Exepcted %s, got: %s", ErrNoVerb, err)
	}
}

func Test_DryRunError(t *testing.T) {
	_, err := New("pip3").DryRun().Seal()
	if err != ErrNoPackageName {
		t.Errorf("Exepcted %s, got: %s", ErrNoPackageName, err)
	}
}

func Test_Params(t *testing.T) {
	cmd, _ := New("pip3").DryRun().Install("pypi").Seal()

	if strings.Compare(*cmd.binaryName, "pip3") != 0 {
		t.Errorf("Not pip3.")
	}
}

func Test_DownloadArgument(t *testing.T) {
	p := &pip{}
	packageName := "examplePackage"
	targetDirectory := "/path/to/target"

	// Call the Download method
	result, _ := p.Download(packageName, targetDirectory).Seal()

	if !contains[string](result.binaryArguments, "-d") {
		t.Errorf("Exepcted %s", "-d")
	}
	if !contains[string](result.binaryArguments, targetDirectory) {
		t.Errorf("Exepcted %s", targetDirectory)
	}
}

func contains[T comparable](s []T, elem T) bool {
	for _, a := range s {
		if a == elem {
			return true
		}
	}
	return false
}
