package lib

import (
    //"os/exec"
    "fmt"
    "os"
)

func StartNewCProject(name string) {
    err := os.Mkdir("~/"+name, 0755)
    if err != nil {
        return
    }

    dir, _ := os.Getwd()
    err = os.Chdir(dir)

    fmt.Println(os.Getwd())


}
