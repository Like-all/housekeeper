package main

import (
    "os"
    "os/exec"
)

func apt(pkg string) () {
    cmd := exec.Command("sudo", "apt-get", "install", pkg)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    cmd.Run()
}
