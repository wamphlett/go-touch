package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing file name to touch")
		os.Exit(1)
	}

	cmd := exec.Command("touch", os.Args[1:]...)
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}