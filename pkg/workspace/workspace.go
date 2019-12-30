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

func Create(folder string) string {
    cwd, err := os.Getwd()
    check(err)
    abs_path := filepath.Join(cwd, folder)
    check(os.Mkdir(abs_path, 0777))
    return abs_path
}

func WriteScript(workdir string, script string) {
    output := []byte(script)
    check(ioutil.WriteFile(filepath.Join(workdir, activatescript), output, 0755))
}
