//go:build !windows

package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"

	"github.com/adrg/xdg"
)

//go:generate wget https://git.savannah.gnu.org/cgit/config.git/plain/config.guess
//go:embed config.guess
var configguess []byte

func main() {
	log.SetFlags(0)
	path, err := xdg.DataFile("go-config/config.guess")
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(path, configguess, 0o755)
	if err != nil {
		log.Fatalln(err)
	}
	// TODO: Use `exec()` to replace own process
	cmd := exec.Command(path, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			log.Fatalln(err)
		}
	}
}
