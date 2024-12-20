package main

import (
	"fmt"
	"os/exec"
)

func copyToClipboard(value string) {
	cmd := exec.Command("wl-copy", value)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
