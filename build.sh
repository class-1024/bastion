#!/bin/bash

workDir=$(pwd)
output="bastion"

export BASTION=${workDir}
export GOPROXY=https://goproxy.io,direct

go version
ls -la
pwd


# 测试构建
go test -v
CGO_ENABLED=0 go build -o $output main.go

# 创建 dist 文件夹
rm -rf dist
mkdir -p dist

cp -r web dist/
cp -r $output dist/
cp -r deploy.sh dist/

ls -la dist

# 删除
rm -rf $output


echo "build success"