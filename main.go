package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	flag.Usage = func() {
		progname := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s [options] command [args ...]\n\n", progname)
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
}

func main() {
	cmd := Command{}
	flag.IntVar(&cmd.Times, "n", 10, "Repeat `number` times")
	flag.DurationVar(&cmd.Duration, "t", 0, "Pause for `interval` times")
	flag.BoolVar(&cmd.Loop, "loop", false, "Infinite loop mode")
	flag.BoolVar(&cmd.Keep, "keep", false, "Keep running if command exit failed")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	cmd.Name = args[0]
	cmd.Args = args[1:]

	cmd.Run()
}
