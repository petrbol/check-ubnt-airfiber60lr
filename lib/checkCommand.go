package lib

import (
	"log"
	"os/exec"
)

func checkCommandAvailable(name []string) {
	for _, i := range name {
		cmd := exec.Command("/bin/sh", "-c", "command -v "+i)
		if err := cmd.Run(); err != nil {
			log.Fatal(err, " application \"", i, "\" not found or not in $PATH")
		}
	}
}
