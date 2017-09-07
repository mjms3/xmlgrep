package main

import (
	"flag"
	"fmt"
	"github.com/mjms3/xmlgrep/extractnodes"
	"io"
	"os"
	"strings"
)

func getReaders(args []string) []io.Reader {

	if len(args) >= 2 {
		var readers []io.Reader
		for _, fileName := range args[1:] {
			reader, err := os.Open(fileName)
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
	subTagToLookFor := flag.String("t", extractnodes.EMPTY_STRING, "Sub tag to filter on.")
	filterToApply := flag.String("f", extractnodes.EMPTY_STRING, "Text filter for Sub tag")
	retainTags := flag.Bool("r", false, "Retain enclosing tags")
	nameSpace := flag.String("n", extractnodes.EMPTY_STRING, "Restrict search to elements in this namespace")

	flag.Parse()

	positionalArgs := flag.Args()
	tagOfInterest := positionalArgs[0]

	filteringParams := extractnodes.ProgramOptions{*subTagToLookFor, *filterToApply,
		*retainTags, *nameSpace}
	readers := getReaders(positionalArgs)
	for _, reader := range readers {
		extractedNodes := extractnodes.ExtractNodes(reader, tagOfInterest, filteringParams)
		for _, node := range extractedNodes {
			trimmedNode := strings.TrimSpace(node)
			if len(trimmedNode) > 0 {
				fmt.Printf("%s\n", trimmedNode)
			}
		}
	}
}
