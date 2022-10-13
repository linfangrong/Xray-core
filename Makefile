export GO111MODULE=on
export GOPROXY=https://goproxy.cn

default:
	go build -o $(shell go env GOPATH)/bin/xray -trimpath -ldflags "-s -w -buildid=" ./main

mips:
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o $(shell go env GOPATH)/bin/xray-linux-mips-softfloat -trimpath -ldflags "-s -w -buildid=" ./main
	upx --best --lzma $(shell go env GOPATH)/bin/xray-linux-mips-softfloat

proto:
	powerproto build proxy/shadowsocks/config.proto
