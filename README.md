# genv is a Go imitation of Python's venv module

## Why?
Go programmers typically put all their code into one workspace. Over time, it gets really difficult to keep track of dependencies and build artifacts (at least, this is the case for me). In Python development, we keep python versions and dependencies isolated into one workspace per project. This helps isolate the dependencies for that project. Go has the additional headache of cleaning up old binaries. This project is meant to create a new workspace for each project.

## How is it supposed to work?
```
$ go get genv
$ genv <new dir>
$ cd <new dir>
$ source activate
```

## Is this method of development endorsed by the Golang community?
No. You are free to use this for your own development. However, as of this writing, the community isn't developing in this way. 
