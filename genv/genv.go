// genv creates a workspace with its own go binary and gopath.
// It's meant to keep track of your go code and its dependencies
// including the go binary used.

package main

import (
        "fmt"
        "os"
        "github.com/nishakm/genv/pkg/script"
        "github.com/nishakm/genv/pkg/workspace"
)

func main() {
    folder := os.Args[1]
    // create the folder
    pathtoactivate := workspace.Create(folder)
    // write the script
    workspace.WriteScript(pathtoactivate, script.Generate(pathtoactivate))
    fmt.Println("Workspace created at %s", pathtoactivate)
}
