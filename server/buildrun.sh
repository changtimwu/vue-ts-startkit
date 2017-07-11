#!/bin/sh
go-bindata dist/...
go build
./server
