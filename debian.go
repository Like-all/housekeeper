package main

import (
    "os"
    "os/exec"
)

func apt(pkg string) () {
    cmd := exec.Command("sudo", "apt-get", "-y", "--no-install-recommends", "install", pkg)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    cmd.Run()
}
