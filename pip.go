package gopip

import (
	"errors"
	"fmt"

	"github.com/CREDOProject/go-pip/shell"
)

var execCommander = shell.NewExecShim

var (
	ErrNoPackageName = errors.New("Package name not specified")
)

type Verb string
type command string

const (
	Install Verb = "install"
	NoVerb       = ""
)

type pip struct {
	binaryName  *string
	dryRun      bool
	verb        Verb
	packageName *string
}

// Start a new Pip command.
// https://pip.pypa.io/en/stable/
func New(binaryName string) *pip {
	return &pip{
		binaryName: &binaryName,
		verb:       NoVerb,
	}
}

// Enable dry-run.
// https://pip.pypa.io/en/stable/cli/pip_install/#cmdoption-dry-run
func (p *pip) DryRun() *pip {
	p.dryRun = true
	return p
}

// Install a package.
// https://pip.pypa.io/en/stable/cli/pip_install/
func (p *pip) Install(packageName string) *pip {
	p.packageName = &packageName
	p.verb = Install
	return p
}

// Seals the command so it can be run.
func (p *pip) Seal() (command, error) {
	if p.packageName == nil && (p.dryRun || p.verb == Install) {
		return "", ErrNoPackageName
	}

	dryRun := ""
	if p.dryRun {
		dryRun = "--dry-run"
	}

	packageName := ""
	if p.packageName != nil {
		packageName = *p.packageName
	}

	templatedCmd := fmt.Sprintf(
		"%s %s %s %s",
		*p.binaryName,
		string(p.verb),
		packageName,
		dryRun,
	)

	return command(templatedCmd), nil
}

// Runs the command.
func (c *command) Run() error {
	error := execCommander().Command(string(*c)).Run()
	return error
}
