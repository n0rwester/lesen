// Package files implements functions for reading and parsing directory data
package files

import (
	"os"
)

type Entry struct {
	Name string
	Type os.FileMode
}

func ReadDir(dir string) []Entry {
	entriesRaw, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	dirs := []Entry{}
	files := []Entry{}

	for _, entry := range entriesRaw {
		if entry.Type().IsDir() {
			dirs = append(dirs, Entry{
				entry.Name(),
				entry.Type(),
			})
			continue
		}

		files = append(files, Entry{
			entry.Name(),
			entry.Type(),
		})
	}

	entries := append(dirs, files...)

	return entries
}
