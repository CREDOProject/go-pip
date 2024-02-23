package gopip

import (
	"errors"
	"fmt"
	"os"

	"github.com/CREDOProject/go-pip/shell"
)

var execCommander = shell.NewExecShim

var (
	ErrNoPackageName = errors.New("Package name not specified.")
	ErrNoVerb        = errors.New("Verb not specified.")
)

type verb string
type command string

const (
	Install verb = "install"
	NoVerb       = ""
)

type pip struct {
	binaryName  *string
	dryRun      bool
	verb        verb
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

	if p.verb == NoVerb {
		return "", ErrNoVerb
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

type RunOptions struct {
	Output *os.File
}

// Runs the command.
func (c *command) Run(options *RunOptions) error {
	command := execCommander().Command(string(*c))
	if options.Output != nil {
		command.Stdout = options.Output
		command.Stderr = options.Output
	}
	error := command.Run()
	return error
}
