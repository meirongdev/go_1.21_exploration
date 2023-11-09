# Go1.21 Exploration

## Compatibility

### Pre-requisites

```shell
gvm use 1.21.1
```

### Change go directive in go.mod

```shell
go get go@1.21.3
```

Go1.21.3 will be save to `~/.gvm/pkgsets/go1.21.1/global/pkg/mod/cache/download/golang.org/toolchain/@v`

### Change go version cmd

Run `go version` in current directory will show `go version go1.21.3 darwin/amd64`, but in other directory will show `go version go1.21.1 darwin/amd64`.

### Available ToolChains

TODO check the code