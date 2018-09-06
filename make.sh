#!/usr/bin/env bash

rm -f tictacgo.zip tictacgo tictacgo.exe

export GOOS="linux"
go build -o tictacgo main.go 

export GOOS="windows"
go build -o tictacgo.exe main.go 

zip tictacgo.zip tictacgo tictacgo.exe
