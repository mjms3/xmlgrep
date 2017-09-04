package main

import (
	"flag"
	"fmt"
	"github.com/mjms3/xmlgrep/extractnodes"
	"io"
	"os"
	"strings"
)

func getReader(args []string) io.Reader {
	if len(args) == 2 {
		fileName := args[1]
		reader, err := os.Open(fileName)
		if err != nil {
			fmt.Errorf("Error opening: %s\n%s", fileName, err)
		}
		return reader
	}
	return os.Stdin
}

func main() {
	subTagToLookFor := flag.String("t", extractnodes.EMPTY_STRING, "Sub tag to filter on.")
	filterToApply := flag.String("f", extractnodes.EMPTY_STRING, "Text filter for Sub tag")
	retainTags := flag.Bool("r", false, "Retain enclosing tags")

	flag.Parse()

	positionalArgs := flag.Args()
	tagOfInterest := positionalArgs[0]
	reader := getReader(positionalArgs)
	filteringParams := extractnodes.ProgramOptions{*subTagToLookFor, *filterToApply,
		*retainTags}
	extractedNodes := extractnodes.ExtractNodes(reader, tagOfInterest, filteringParams)
	for _, node := range extractedNodes {
		trimmedNode := strings.TrimSpace(node)
		if len(trimmedNode) > 0 {
			fmt.Printf("%s\n", trimmedNode)
		}
	}

}
