export GO111MODULE=on
export GOPROXY=https://goproxy.cn

default:
	go build -o $(shell go env GOPATH)/bin/xray -trimpath -ldflags "-s -w -buildid=" ./main

proto:
	powerproto build proxy/shadowsocks/config.proto

