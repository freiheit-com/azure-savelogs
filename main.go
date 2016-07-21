package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var (
		home   = os.Getenv("HOME")
		logDir = path.Join(home, "LogFiles", "application")
		today  = time.Now().Format("200601021504050700")
	)
	os.MkdirAll(logDir, os.ModeDir|os.ModePerm)

	outFile, _ := os.Create(path.Join(logDir, fmt.Sprintf("%s-stdout.log", today)))
	defer outFile.Close()
	errFile, _ := os.Create(path.Join(logDir, fmt.Sprintf("%s-stderr.log", today)))
	defer errFile.Close()

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = outFile
	cmd.Stderr = errFile
	cmd.Run()
}
