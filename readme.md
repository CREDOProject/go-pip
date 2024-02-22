# Go-Pip

Go bindings to manage Pip.

---

Example:

```go
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
```
