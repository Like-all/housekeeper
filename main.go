package main

import (
    "fmt"
    "os"
)

var CfgParams, _ = loadConfig()

func main() {
    recipe, err := getRecipe(CfgParams.Mirror, os.Args[1])
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    pm, err := detectPackageManager()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    for i := range recipe.Preinstall {
        shell(recipe.Preinstall[i])
    }

    for i := range recipe.Packages {
        switch pm {
            case "apt":
                apt(recipe.Packages[i])
        }
    }

    for i := range recipe.Dotfiles {
        getDotfile(CfgParams.Mirror, recipe.Dotfiles[i])
    }

    for i := range recipe.Postinstall {
        shell(recipe.Postinstall[i])
    }
}
