package script

import (
    "fmt"
    
    "github.com/nishakm/genv/pkg/workspace"
)

var script = `# This file is generated by genv
# This file must be used with "source activate" *from bash*
# You cannot run it directly

deactivate () {
  # reset old path
  if [ -n "${OLD_PATH:-}" ] ; then
    PATH="${OLD_PATH:-}"
    unset OLD_PATH
  fi

  # reset old gopath
  if [ -n "${OLD_GOPATH:-}" ] ; then
    GOPATH="${OLD_GOPATH:-}"
    export GOPATH
    unset OLD_GOPATH
  fi

  # Call the hash command for bash and zsh to make it forget past commands
  # or else the new exports may not be respected
  if [ -n "${BASH:-}" -o -n "${ZSH_VERSION:-}" ] ; then
    hash -r
  fi

  # Remove the new prompt
  if [ -n "${OLD_PS1:-}" ] ; then
    PS1="${OLD_PS1:-}"
    export PS1
    unset OLD_PS1
  fi
}

# our new go path is here
NEW_GOPATH=%s

# store current path
OLD_PATH=$PATH
# store current gopath
OLD_GOPATH=$GOPATH

# set new path
PATH="%s:%s:$PATH"
export PATH

# set new gopath
GOPATH="$NEW_GOPATH"
export GOPATH

# set new prompt
OLD_PS1="${PS1:-}"
PS1="(%s) ${PS1:-}"

# Call the hash command for bash and zsh to make it forget past commands
# or else the new exports may not be respected
if [ -n "${BASH:-}" -o -n "${ZSH_VERSION:-}" ] ; then
  hash -r
fi
`
func Generate(gopath string, envpath string, folder string) string {
    binpath := workspace.Binpath(gopath)
    return fmt.Sprintf(script, gopath, binpath, envpath, folder)
}
