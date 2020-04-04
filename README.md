# go with micro

## Install etcd/consul
```bash
brew install etcd
```

## Install protobuf code generation
```bash
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro/v2
```

## Generate Code
```bash
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```

## Run
```bash
etcd
MICRO_REGISTRY=etcd go run main.go
```
