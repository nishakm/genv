// genv creates a workspace with its own go binary and gopath.
// It's meant to keep track of your go code and its dependencies
// including the go binary used.

package main

import (
        "fmt"
        "os"
        "flag"
        "github.com/nishakm/genv/pkg/script"
        "github.com/nishakm/genv/pkg/workspace"
        "github.com/nishakm/genv/pkg/versions"
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [options] folder\n", os.Args[0])
        flag.PrintDefaults()
    }
    verPtr := flag.String("version", "", "Go version")
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
}
