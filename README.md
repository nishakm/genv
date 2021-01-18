# genv is a Go imitation of Python's venv module

## Why?
Go programmers typically put all their code into one workspace. Over time, it gets really difficult to keep track of dependencies and build artifacts (at least, this is the case for me). In Python development, we keep python versions and dependencies isolated into one workspace per project. This helps isolate the dependencies for that project. Go has the additional headache of cleaning up old binaries. This project is meant to create a new workspace for each project.

## How is it supposed to work?
To create a development environment
```
$ go get github.com/nishakm/genv
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
$ rm -rf test
```

## Controlling go versions
Genv assumes alternative go versions are discoverable in your current $PATH. If they are, you can use them in your isolated environment
```
$ genv -version go1.14.4 test
$ cd test
$ source bin/activate
(test)$ go version
go version go1.14.4 linux/amd64
``` 
