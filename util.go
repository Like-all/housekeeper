package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "os"
    "os/user"
    "bufio"
    "io/ioutil"
    "bytes"
    "runtime"
)

type Config struct {
    Mirror string
}

type Recipe struct {
    Preinstall []string
    Packages []string
    Dotfiles []string
    Postinstall []string
}

func initialize() () {
    usr, _ := user.Current()
    os.MkdirAll(usr.HomeDir + "/.config/housekeeper/", 0755)

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Please specify recipes repository URL:")
    scanner.Scan()
    mirror := scanner.Text()

    configFile, err := os.Create(usr.HomeDir + "/.config/housekeeper/config.json")
    if err != nil {
        fmt.Println(err.Error())
    }


    _, err = configFile.WriteString("{\n    \"mirror\": \"" + mirror + "\"\n}\n")
    if err != nil {
        fmt.Println(err.Error())
    }
}

func loadConfig() (c *Config, err error) {
    var bfile []byte
    usr, _ := user.Current()
    cfgpath := usr.HomeDir + "/.config/housekeeper/config.json"
    if bfile, err = ioutil.ReadFile(cfgpath); err != nil {
        initialize()
        loadConfig()
    }
    c = new(Config)
    err = json.Unmarshal(bfile, c)
    return
}

func getRecipe(mirror, recipeName string) (recipe *Recipe, err error) {
    resp, err := http.Get(mirror + "/" + recipeName + ".json")
    if err != nil {
        return nil, err
    }
    if resp.StatusCode == 404 {
        return nil, fmt.Errorf("Recipe not found.\n")
    }
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("Unexpected status code: %d\n", resp.StatusCode)
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    recipeJson := buf.String()

    recipe = &Recipe{}
    json.Unmarshal([]byte(recipeJson), &recipe)

    return
}

func detectPackageManager() (pm string, err error){
    switch runtime.GOOS {
        case "linux":
            if _, err := os.Stat("/usr/bin/apt-get"); err == nil {
                pm = "apt"
                return pm, nil
            }
            if _, err := os.Stat("/usr/bin/yum"); err == nil {
                pm = "yum"
                return pm, nil
            }
            err = fmt.Errorf("Looks like you running housekeeper on LFS\n")
    }
    err = fmt.Errorf("Looks like you running housekeeper on unsupported platform\n")
    return
}
