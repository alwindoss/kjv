package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed data
var kjvFS embed.FS

func main() {
	f, err := kjvFS.Open("data/kjv-1769.txt")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}