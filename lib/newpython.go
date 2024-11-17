package lib

import (
    "os/exec"
    "fmt"
    "os"
)

func StartNewPythonProject(name string) {
 
    //path := "~/"+name 
    cmd := exec.Command("mkdir", "~/"+name)
    err := cmd.Run()

    if err != nil {
        fmt.Println("Error on making Project Folder!", err)
        return
    } else {
        fmt.Printf("The directory %s was created...\n", name)
    }

    err = os.Chdir(name)
    
    if err != nil {
        fmt.Println("Error, please try again.")
        return 
    } else {

        file, err := os.Create("main.py")
        
        if err != nil {
            fmt.Println("Error, please try again.")
            return 
        }
        
        _, err = file.WriteString("if __name__ == '__main__':\n  print('all is good!')")
        if err != nil {
            fmt.Println("Error, please try again.")
        }

        srcpath := "~/%s/src"+name
        libpath := "~/%s/lib"+name
        os.Mkdir(srcpath, 0755)
        os.Mkdir(libpath, 0755)

        cmd := exec.Command("python3", "-m", "venv", "venv")
        err = cmd.Run()
        if err != nil {
            fmt.Println("Error! Possibly, python3 is not installed.")
            return 
        }

        venv := "~/%s/venv/bin/activate"
        exec.Command("source", venv)

        fmt.Print("\nInit Git? [y/n] ")
        var gitOrNo string
        fmt.Scan(&gitOrNo)
        
        if gitOrNo == "y" || gitOrNo == "Y" || gitOrNo == "Yes" || gitOrNo == "yes" {
            err := exec.Command("git", "init")
            if err != nil {
                fmt.Println("Error! Possibly, git is not installed.")
                return
            }
            exec.Command("touch", ".gitignore")
        } else {
            fmt.Print("\n")
        }
    } 
    os.Chdir("~/")
    fmt.Println("Done!")
}
