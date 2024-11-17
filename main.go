package main

import (
    "fmt"
    "os"
    "start/lib"
)

var (
    LANG = os.Args[1]
    NAME = os.Args[2]
)

func main() {
    if len(os.Args) == 3 {
        if LANG == "python" || LANG == "py" {
            lib.StartNewPythonProject(NAME)
        } else if LANG == "c" || LANG == "C" {
            lib.StartNewCProject(NAME)
        } else {
            fmt.Println("Language not found!")
        }
    } else {
        fmt.Println("Too many args.")
        return 
    }

}
