package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

type Command struct {
	Name     string
	Args     []string
	Times    int
	Duration time.Duration
	Loop     bool
	Keep     bool
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
			c.sleep()
		}
	} else {
		for i := 0; i < c.Times; i++ {
			c.exec()
			c.sleep()
		}
	}
}

func (c *Command) sleep() {
	if c.Duration > 0 {
		time.Sleep(c.Duration)
	}
}
