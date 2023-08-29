# Govm
Manage Go versions on your local machine and seemingly test how your applications run on different versions

## Quickstart guide
### Install new version
Install a specific version of Go using the semantic versioning scheme
```
> govm install 1.21.0
> govm install 1.21
> govm install 1
```
or grab the latest version by passing in "go"
```
> govm install go
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