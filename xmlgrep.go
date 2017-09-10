package main

import (
	"flag"
	"fmt"
	"github.com/mjms3/xmlgrep/extractnodes"
	"io"
	"os"
	"strings"
)

type fileSystem interface {
	Open(name string) (io.Reader, error)
}

type osFS struct{}

func (osFS) Open(name string) (io.Reader, error) { return os.Open(name) }

func getReaders(fs fileSystem, args []string) []io.Reader {
	if len(args) >= 2 {
		var readers []io.Reader
		for _, fileName := range args[1:] {
			reader, err := fs.Open(fileName)
			if err != nil {
				fmt.Errorf("Error opening: %s\n%s", fileName, err)
			}
			readers = append(readers, reader)
		}
		return readers
	}

	return []io.Reader{os.Stdin}
}

func main() {
	var fs fileSystem = osFS{}
	output := os.Stdout
	run(fs, output)
}
func run(fs fileSystem, output io.Writer) {
	subTagToLookFor := flag.String("t", extractnodes.EMPTY_STRING, "Sub tag to filter on.")
	filterToApply := flag.String("f", extractnodes.EMPTY_STRING, "Text filter for Sub tag")
	retainTags := flag.Bool("r", false, "Retain enclosing tags")
	nameSpace := flag.String("n", extractnodes.EMPTY_STRING, "Restrict search to elements in this namespace")
	flag.Parse()
	positionalArgs := flag.Args()
	tagOfInterest := positionalArgs[0]
	filteringParams := extractnodes.ProgramOptions{*subTagToLookFor, *filterToApply,
		*retainTags, *nameSpace}
	readers := getReaders(fs, positionalArgs)
	for _, reader := range readers {
		extractedNodes := extractnodes.ExtractNodes(reader, tagOfInterest, filteringParams)
		for _, node := range extractedNodes {
			trimmedNode := strings.TrimSpace(node)
			if len(trimmedNode) > 0 {
				fmt.Fprintf(output, "%s\n", trimmedNode)
			}
		}
	}
}
