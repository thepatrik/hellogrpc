# proto

gRPC skeleton for golang.

### setup

Download the latest release of Protocol Buffers from [here](https://developers.google.com/protocol-buffers/docs/downloads).

Get protoc-gen-go:
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
This will download the binary to `$GOPATH/bin`. Make sure that it is in the `$PATH`
```
export PATH="$GOPATH/bin:$PATH"
```

To generate protobuf stubs run the default make goal:

```
make [gen]
```

The generated files will be placed in the `build` directory.
