package script

import "fmt"

var script = `deactivate () {
  if [ -n "${OLD_GOPATH:-}" ] ; then
    GOPATH="${OLD_GOPATH:-}"
    export GOPATH
  fi
}

# our new go path is the current working directory
NEW_GOPATH=%s

# store current gopath
OLD_GOPATH=$GOPATH

# set new gopath
GOPATH="$NEW_GOPATH"
export GOPATH
`
func Generate(workspace string) string {
    return fmt.Sprintf(script, workspace)
}
