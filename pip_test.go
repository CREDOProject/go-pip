package gopip

import (
	"testing"

	"github.com/CREDOProject/go-pip/utils"
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

	pip, err := utils.DetectPipBinary()
	if err != nil {
		// TODO: No pip binary in system
	}
	command, err := New(pip).Install("conda").DryRun().Seal()
	if err != nil {
		// TODO: Error building command
	}
	err = command.Run()
	if err != nil {
		// TODO: Error running command
	}
}
