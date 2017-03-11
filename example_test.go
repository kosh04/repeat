package main

import "runtime"

func Example() {
	cmd := Command{
		Name:  "echo",
		Args:  []string{"HELLO"},
		Times: 5,
	}
	if runtime.GOOS == "windows" {
		cmd.Name = "/msys64/usr/bin/echo.exe"
		//cmd.Name = "cmd.exe"
		//cmd.Args = []string{"/c", "echo", "HELLO"}
	}

	cmd.Run()
	// Output:
	// HELLO
	// HELLO
	// HELLO
	// HELLO
	// HELLO
}
