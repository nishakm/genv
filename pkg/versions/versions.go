package versions

import (
    "fmt"
    "os/exec"
    "strings"
)

func GetGoRoot(go_version string) string {
    result, err := exec.Command("which", go_version).Output()
    if err != nil {
        fmt.Printf("Error in finding executable for %s: %s", go_version, err)
    }
    return strings.TrimRight(string(result), "\n")
}
