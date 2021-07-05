// Package project provides filepaths for corresponding remote VCSs
// Currently, github.com is the only one implemented
package project

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

var githubSsh string = "git@"
var githubHttps string = "https://"
var githubSuffix string = ".git"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Clone execs "git clone project local/path" according
// to Go's expectation of where the project should be
// in the gopath.
func CloneProject(proj, ppath string) {
    // make the project path first
    check(os.MkdirAll(ppath, 0777))
    
    // NOTE: Authentication is not handled here
    cmd := exec.Command("git", "clone", proj, ppath)
    if err := cmd.Run(); err != nil {
        fmt.Println("Cannot clone project %s: %s\n", proj, err)
    }
}

// GetProjPath extracts the project path from the given string.
// Typically for Github, the URI looks like
// git@github.com:path/to/project.git or
// https://github.com/path/to/project.git
// We return the "github.com/path/to/project" bit
func GetProjPath(proj string) string {
    if strings.Contains(proj, githubSsh) {
        domainPath := strings.Split(proj, githubSsh)[1]
        slashPath := strings.Replace(domainPath, ":", "/", 1)
        return strings.Split(slashPath, githubSuffix)[0]
    }
    if strings.Contains(proj, githubHttps) {
        slashPath := strings.Split(proj, githubHttps)[1]
        return strings.Split(slashPath, githubSuffix)[0]
    }
    // if the parsing didn't work, just return nothing
    return ""
}
