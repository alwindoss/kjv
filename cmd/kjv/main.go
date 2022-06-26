package main

import (
	"fmt"
	"io"

	"github.com/alwindoss/kjv"
)

func main() {
	f, err := kjv.FS.Open("kjvsrc/kjv-1769.txt")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}