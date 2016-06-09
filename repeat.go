package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	times = flag.Int("n", 10, "Repeat N times")
	loop  = flag.Bool("loop", false, "Infinite loop mode")
	keep  = flag.Bool("keep", false, "Keep running if command exit failed")
)

func init() {
	flag.Usage = func() {
		name := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s [options] command [args ...]\n\n", name)
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *loop {
		for {
			executeCommand(args[0], args[1:]...)
		}
	} else {
		for i := 0; i < *times; i++ {
			executeCommand(args[0], args[1:]...)
		}
	}
}

func executeCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil && *keep != true {
		log.Fatal(err)
	}
}
