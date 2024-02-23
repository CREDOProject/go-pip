package gopip

import (
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
