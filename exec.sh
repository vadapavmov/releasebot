#!/bin/zsh

go build -o bot main.go structs.go  discord.go env.go utils.go

./bot -guild 1184866178359885864
