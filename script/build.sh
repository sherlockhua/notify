#!/bin/bash

MODULE_NAME=notify
SOURCE_DIR=
if [ ! -e ./go.mod ];then
    go mod init
fi

go mod tidy

go build -o output/bin/$MODULE_NAME $SOURCE_DIR
cp -r configs output/
cp ./script/start.sh ./output/
chmod +x ./output/start.sh