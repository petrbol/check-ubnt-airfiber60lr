package lib

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func parseArgs(args []string) (*ubntCheckStartupOpts, error) {
	opts := &ubntCheckStartupOpts{}
	_, err := flags.ParseArgs(opts, args)
	return opts, err
}

func Start() {
	opts, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	opts.getDeviceData()
}
