package gopip

import (
	"errors"
	"os"

	"github.com/CREDOProject/go-pip/shell"
)

var execCommander = shell.NewExecShim

var (
	ErrNoPackageName = errors.New("Package name not specified.")
	ErrNoVerb        = errors.New("Verb not specified.")
)

type verb string
type command struct {
	binaryName      *string
	binaryArguments []string
}

const (
	Install  verb = "install"
	Download verb = "download"
	NoVerb        = ""
)

type pip struct {
	binaryName      *string
	dryRun          bool
	verb            verb
	packageName     *string
	targetDirectory *string
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

// Downloads a package.
// https://pip.pypa.io/en/stable/cli/pip_download/
func (p *pip) Download(packageName string, targetDirectory string) *pip {
	p.packageName = &packageName
	p.targetDirectory = &targetDirectory
	p.verb = Download
	return p
}

// Seals the command so it can be run.
func (p *pip) Seal() (command, error) {
	if p.packageName == nil && (p.dryRun || p.verb == Install) {
		return command{}, ErrNoPackageName
	}

	arguments := []string{}

	if p.verb == NoVerb {
		return command{}, ErrNoVerb
	}

	arguments = append(arguments, string(p.verb))

	if p.packageName != nil {
		arguments = append(arguments, *p.packageName)
	}

	if p.dryRun {
		arguments = append(arguments, "--dry-run")
	}

	if p.targetDirectory != nil {
		arguments = append(arguments, "-d", *p.targetDirectory)
	}

	return command{
		binaryName:      p.binaryName,
		binaryArguments: arguments,
	}, nil
}

type RunOptions struct {
	Output *os.File
	Env    []string
}

// Runs the command.
func (c *command) Run(options *RunOptions) error {
	command := execCommander().Command(*c.binaryName, c.binaryArguments...)
	if options.Output != nil {
		command.Stdout = options.Output
		command.Stderr = options.Output
	} else {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}
	if options.Output != nil {
		command.Env = options.Env
	}
	error := command.Run()
	return error
}
