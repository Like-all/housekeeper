package main

import (
    "bytes"
    "os"
    "os/user"
    "fmt"
    "path/filepath"
    "net/http"
)

func getDotfile(mirror, dotfilePath string) () {
    resp, err := http.Get(mirror + "/dotfiles/" + dotfilePath)
    if err != nil {
        fmt.Println(err.Error())
    }
    if resp.StatusCode == 404 {
        fmt.Println("Dotfile " + dotfilePath + " not found in the repository.")
    }
    if resp.StatusCode != 200 {
        fmt.Printf("Unexpected status code: %d", resp.StatusCode)
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    dotfileBody := buf.String()

    usr, _ := user.Current()
    os.MkdirAll(filepath.Dir(usr.HomeDir + "/" + dotfilePath), 0700)

    dotfile, err := os.Create(usr.HomeDir + "/" + dotfilePath)
    if err != nil {
        fmt.Println(err.Error())
    }

    _, err = dotfile.WriteString(dotfileBody)
    if err != nil {
        fmt.Println(err.Error())
    }
}
