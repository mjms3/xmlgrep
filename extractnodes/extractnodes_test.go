package extractnodes

import (
	"testing"
	"reflect"
	"strings"
)

func TestExtractNodes(t *testing.T){
	inputString := `<data><A>Test1</A><A>Test2</A></data>`
	inString := strings.NewReader(inputString)
	outArray := []string{`Test1`,`Test2`}
	extractedNodes := ExtractNodes(inString,"A")
	if !reflect.DeepEqual(outArray,extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inputString, extractedNodes, outArray)
	}

}
