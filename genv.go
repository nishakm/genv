// genv creates a workspace with its own go binary and gopath.
// It's meant to keep track of your go code and its dependencies
// including the go binary used.

package main

import (
        "fmt"
        "io/ioutil"
        "os"
        "os/exec"
        "flag"
        "path/filepath"

        "github.com/nishakm/genv/pkg/script"
        "github.com/nishakm/genv/pkg/workspace"
        "github.com/nishakm/genv/pkg/versions"
        "github.com/nishakm/genv/pkg/project"
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [options] folder\n", os.Args[0])
        flag.PrintDefaults()
    }
    verPtr := flag.String("version", "", "Provide a Go version to use. Eg: go1.16.5")
    projPtr := flag.String("project", "", "Provide a Git project to clone. Eg: git@github.com/nishakm/genv")
    noTools := flag.Bool("no-tools", false, "Set the no-tools flag to avoid installing envtool in your environment")
    flag.Parse()
    folder := flag.Arg(0)
    // create the workspace with the folder
    folderpath := workspace.Create(folder)
    // get the location of the new gopath
    gopath := workspace.Gopath(folderpath)
    // get the location of the new environment path
    envpath := workspace.Envpath(folderpath)
    // write the script
    workspace.WriteScript(envpath, script.Generate(gopath, envpath, folder))
    fmt.Printf("Workspace created at %s\n", folderpath)
    // if a version is provided, make a symlink
    if *verPtr != "" {
        goroot := versions.GetGoRoot(*verPtr)
        if goroot != "" {
            workspace.SetGoSym(goroot, envpath)
        }
    }
    // if a git project is provided, clone it in the required location
    if *projPtr != "" {
        srcpath := workspace.Srcpath(gopath)
        projPath := filepath.Join(srcpath, project.GetProjPath(*projPtr))
        project.CloneProject(*projPtr, projPath)
    }
    // if we haven't explicitly said to NOT install envtool, we attempt to install it
    if *noTools == false {
        // install the envtool binary
        if *verPtr != "" {
            err := exec.Command(*verPtr, "install", "github.com/nishakm/genv/envtool@latest").Run()
            if err != nil {
                fmt.Println("Error installing envtool: %s", err)
                os.Exit(1)
            }
        } else {
            err := exec.Command("go", "install", "github.com/nishakm/genv/envtool@latest").Run()
            if err != nil {
                fmt.Println("Error installing envtool: %s", err)
                os.Exit(1)
            }
        }
        // find the path the binary is located
        result, err := exec.Command("which", "envtool").Output()
        if err != nil {
            fmt.Println("Error finding envtool: %s", err)
            os.Exit(1)
        }
        envtoolPath := filepath.Join(workspace.Binpath(gopath), "envtool")
        // copy binary over to the environment
        copyFile(string(result), envtoolPath)
        fmt.Println("envtool installed")
    }
}

// Go doesn't have a builtin copy function to copy files
// from one place to another. So I'll make my own
func copyFile(srcPath string, destPath string) {
    content, err := ioutil.ReadFile(srcPath)
    if err != nil {
        fmt.Println("Error reading envtool binary: %s", err)
        os.Exit(1)
    }
    writeErr := ioutil.WriteFile(destPath, content, 0755)
    if err != nil {
        fmt.Println("Error writing envtool to env path: %s", writeErr)
        os.Exit(1)
    }
    

}
