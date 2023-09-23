# Govm
Manage Go versions on your local machine and seemingly test how your applications run on different versions

## Installation
### Linux
1. Download the latest [release](https://github.com/dkartachov/govm/releases) and store it in `$HOME/.govm/bin`
```
mkdir -p $HOME/.govm/bin && curl -L github.com/dkartachov/govm/releases/download/v1.0.0/govm -o $HOME/.govm/bin/govm
```
2. Add `$HOME/.govm/bin` to the `PATH` environment variable:
```
export PATH=$PATH:$HOME/.govm/bin
```

## Quickstart guide
### Install new version
Install the latest version of Go
```
> govm install go
```
or install a specific version using semantic versioning
```
> govm install 1.21.0
> govm install 1.21
> govm install 1
```

### Change default version
You can change the default version used when running the `go` command:
```
> go version
go version go1.20.7 linux/amd64

> govm use 1.21.0
> go version
go version go1.21.0 linux/amd64
```

### Run applications using other versions
You can run applications using previously installed versions without having to change the default version:
```
> go run hello.go
Hello, World!

> go1.20.7 run hello.go
Hello, World!
```