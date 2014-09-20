package main

import (
    "os"
    "os/exec"
)

func shell(action string) (){
    cmd := exec.Command("sh", "-c", action)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    cmd.Run()
}
