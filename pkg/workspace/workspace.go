package workspace

import (
        "os"
        "path/filepath"
        "io/ioutil"
)

var activatescript = "activate"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Gopath(folderpath string) string {
    gopath := filepath.Join(folderpath, "gopath")
    return gopath
}

func Create(folder string) string {
    cwd, err := os.Getwd()
    check(err)
    abs_path := filepath.Join(cwd, folder)
    go_src := filepath.Join(abs_path, "gopath", "src")
    go_bin := filepath.Join(abs_path, "gopath", "bin")
    check(os.MkdirAll(go_src, 0777))
    check(os.MkdirAll(go_bin, 0777))
    return abs_path
}

func WriteScript(workdir string, script string) {
    output := []byte(script)
    check(ioutil.WriteFile(filepath.Join(workdir, activatescript), output, 0755))
}
