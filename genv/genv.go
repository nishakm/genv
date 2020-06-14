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
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [options] folder\n", os.Args[0])
        flag.PrintDefaults()
    }
    verPtr := flag.String("version", "", "Go version")
    flag.Parse()
    folder := flag.Arg(0)
    fmt.Println("version", *verPtr)
    // create the workspace with the folder
    pathtoactivate := workspace.Create(folder)
    // get the location of the new gopath
    gopath := workspace.Gopath(pathtoactivate)
    // write the script
    workspace.WriteScript(pathtoactivate, script.Generate(gopath, folder))
    fmt.Printf("Workspace created at %s\n", pathtoactivate)
}
