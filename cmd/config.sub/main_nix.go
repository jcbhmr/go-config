//go:build aix || android || darwin || dragonfly || freebsd || linux || netbsd || openbsd || plan9 || solaris

package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"

	"github.com/adrg/xdg"
)

// TODO: Use a portable Go-integrated version of Wget

//go:generate wget -q https://git.savannah.gnu.org/cgit/config.git/plain/config.sub -O config.sub
//go:embed config.sub
var configsub []byte

func main() {
	log.SetFlags(0)
	path, err := xdg.DataFile("go-config/config.sub")
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(path, configsub, 0o755)
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
