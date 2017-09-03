package main

import (
	"flag"
	"fmt"
	"github.com/mjms3/xmlgrep/extractnodes"
	"io"
	"os"
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

	flag.Parse()

	positionalArgs := flag.Args()
	tagOfInterest := positionalArgs[0]
	reader := getReader(positionalArgs)
	filteringParams := extractnodes.FilteringParams{*subTagToLookFor, *filterToApply}
	extractedNodes := extractnodes.ExtractNodes(reader, tagOfInterest, filteringParams)
	for _, node := range extractedNodes {
		fmt.Printf("%s\n", node)
	}

}
