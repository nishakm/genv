# genv is a Go imitation of Python's venv module

## Why?
Go programmers typically put all their code into one workspace. Over time, it gets really difficult to keep track of dependencies and build artifacts (at least, this is the case for me). In Python development, we keep python versions and dependencies in one workspace per project. This helps isolate the dependencies for that project. Go has the additional headache of cleaning up old binaries. This project is meant to create a new workspace for each project.

## How to use it
To create a development environment
```
$ go install github.com/nishakm/genv@latest
$ genv test
$ cd test
$ source bin/activate
```
To undo the dev environment
```
$ deactivate
```
To remove the dev environment
```
$ sudo rm -rf test
```
(note: root permissions are required as go sets all the file permissions to read only)

## Controlling go versions
You can create dev environments for specific go versions. If the version doesn't exist in your dev environment, genv will download it for you.
```
$ genv -version go1.14.4 test
$ cd test
$ source bin/activate
(test)$ go version
go version go1.14.4 linux/amd64
``` 

## Cloning a Go project
If you want to seed the dev environment with a specific Go project, you can do the following:
```
$ genv -project git@github.com/nishakm/genv test
$ cd test
$ source bin/activate
(test)$ cd gopath/src/github.com/nishakm/genv
```

Currently GitHub cloning via SSH and HTTPS are supported.
