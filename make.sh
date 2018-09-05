#!/usr/bin/env bash

export GOOS="linux"
go build -o tictacgo main.go board.go position.go turn.go

export GOOS="windows"
go build -o tictacgo.exe main.go board.go position.go turn.go
