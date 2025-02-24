package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: sponge <filename>\n")
		os.Exit(1)
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.WriteString("Error reading from stdin: " + err.Error() + "\n")
		os.Exit(1)
	}

	filename := os.Args[1]

	fileInfo, err := os.Stat(filename)
	if err != nil {
		os.Stderr.WriteString("Error getting file info: " + err.Error() + "\n")
		os.Exit(1)
	}

	err = os.WriteFile(filename, input, fileInfo.Mode())
	if err != nil {
		os.Stderr.WriteString("Error writing to file: " + err.Error() + "\n")
		os.Exit(1)
	}
}
