package main

import (
	"os"

	"github.com/n0rwester/lesen/files"
	"github.com/n0rwester/lesen/render"
)

func main() {
	if len(os.Args) == 1 {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		entries := files.ReadDir(wd)
		render.RenderEntries(entries, wd)
	}
}
