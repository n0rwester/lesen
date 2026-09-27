// Package render implements functions for prettyprinting directory entries
package render

import (
	"fmt"
	"strings"

	"github.com/n0rwester/lesen/files"
	"github.com/n0rwester/lesen/util"
)

func RenderEntries(dir []files.Entry, wd string) {
	entries := []string{}
	for _, entry := range dir {
		if entry.Type.IsDir() {
			entries = append(entries, renderDir(entry.Name))
		} else if entry.Type.IsRegular() {
			entries = append(entries, renderFile(entry.Name))
		}
	}

	fmt.Println("") // upper padding
	for _, line := range box(entries, wd) {
		fmt.Println(line)
	}
	fmt.Println("") // lower padding
}

func renderDir(dir string) string {
	return " " + dir + "/"
}

func renderFile(file string) string {
	return "󰈔 " + file
}

func box(lines []string, name string) []string {
	output := []string{}

	maxlen := util.GetMaxLen(lines)
	maxlen = max(len(name)-1, maxlen, 7)

	output = append(output, " ┌─"+util.Pad(name, maxlen, "─")+"─┐")
	for _, line := range lines {
		// Accounting for nerd font icons which tend to have
		// inconsistent widths
		icon := strings.Split(line, " ")
		iconLength := len(icon[0]) - 1

		output = append(output, " │ "+util.Pad(line, maxlen+iconLength+1, " ")+"│")
	}
	output = append(output, " └─"+util.Pad("", maxlen, "─")+"─┘")
	return output
}
