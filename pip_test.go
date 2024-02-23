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
