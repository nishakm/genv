package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"

    "github.com/nishakm/genv/pkg/versions"
    "github.com/nishakm/genv/pkg/workspace"
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s go_version\n", os.Args[0])
        flag.PrintDefaults()
    }
    flag.Parse()
    goVersion := flag.Arg(0)
    // We assume that the environment has been activated
    // so we find the current GOPATH
    gopath := os.Getenv("GOPATH")
    // our environment path is relative to the gopath
    folderpath := filepath.Dir(gopath)
    // now we can get the other paths
    envpath := workspace.Envpath(folderpath)
    // get the path of the new go version
    newGoRoot := versions.GetGoRoot(goVersion)
        if newGoRoot != "" {
            // we have a new go binary to symlink to
            workspace.SetGoSym(newGoRoot, envpath)
        }
        fmt.Printf("Go version changed to %s\n. Run 'deactivate' and 'source bin/activate' again", goVersion)
}
