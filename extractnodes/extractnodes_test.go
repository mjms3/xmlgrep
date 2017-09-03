package extractnodes

import (
	"testing"
	"reflect"
)

func TestExtractNodes(t *testing.T){
	inString := `<data><A>Test1</A><A>Test2</A></data>`
	outArray := []string{`Test1`,`Test2`}
	extractedNodes := ExtractNodes(inString)
	if !reflect.DeepEqual(outArray,extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inString, extractedNodes, outArray)
	}

}
