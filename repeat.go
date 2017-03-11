package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Command struct {
	Name  string
	Args  []string
	Times int
	Loop  bool
	Keep  bool
}

func (c *Command) exec() {
	cmd := exec.Command(c.Name, c.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil && !c.Keep {
		log.Fatal(err)
	}
}

func (c *Command) Run() {
	if c.Loop {
		for {
			c.exec()
		}
	} else {
		for i := 0; i < c.Times; i++ {
			c.exec()
		}
	}
}

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
