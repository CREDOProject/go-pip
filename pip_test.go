package gopip

import (
	"testing"
)

func Test_New(t *testing.T) {
	_, err := New("pip3").Seal()
	if err != nil {
		t.Error(err)
	}
}

func Test_DryRunError(t *testing.T) {
	_, err := New("pip3").DryRun().Seal()
	if err != ErrNoPackageName {
		t.Errorf("Exepcted %s, got: %s", ErrNoPackageName, err)
	}
}
