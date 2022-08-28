package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const code = `package main

func main() {
	panic("oops")
}
`

func runSilent(pwd string, args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = pwd
	if out, err := cmd.CombinedOutput(); err != nil {
		os.Stdout.Write(out)
		panic(err)
	}
}

func main() {
	root, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := os.RemoveAll(root); err != nil {
			panic(err)
		}
	}()
	fmt.Printf("Running within %s\n", root)

	gopkg := filepath.Join(root, "inner")
	if err = os.Mkdir(gopkg, 0o700); err != nil {
		panic(err)
	}
	runSilent(gopkg, "go", "mod", "init", "inner")
	if err = os.WriteFile(filepath.Join(gopkg, "main.go"), []byte(code), 0o600); err != nil {
		panic(err)
	}
	exe := filepath.Join(root, "bin.exe")
	runSilent(gopkg, "go", "build", "-o", exe, ".")
	cmd := exec.Command(exe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = root
	cmd.Run()
}
