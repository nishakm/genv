package versions

import (
    "fmt"
    "os/exec"
    "strings"
)

func GetGoRoot(go_version string) string {
    result, err := exec.Command("which", go_version).Output()
    if err != nil {
        fmt.Printf("Getting go version: %s\n", go_version)
        dlerr := getGoVersion(go_version)
        if dlerr != nil {
            return ""
        }
        // try again
        result, err = exec.Command("which", go_version).Output()
        if err != nil {
            fmt.Printf("╮(╯ _╰ )╭")
            return ""
        }
    }
    return strings.TrimRight(string(result), "\n")
}

// internal function to download go version that
// doesn't exist on the system
// NOTE: go versions are of the form go<semantic_version>
// if the full string doesn't exist remotely from
// golang.org/dl then this function will return nil 
func getGoVersion(go_version string) error {
    golang_remote := fmt.Sprintf("golang.org/dl/%s", go_version)
    getErr := exec.Command("go", "get", golang_remote).Run()
    if getErr != nil {
        fmt.Printf("Error in finding go version %s on golang.org/dl\n", go_version)
        return getErr
    }
    dlErr := exec.Command(go_version, "download").Run()
    if dlErr != nil {
        // there may be an error because the version is already downloaded
        fmt.Printf("Download error: %s", dlErr)
    }
    return nil
}
