package extractnodes

import "testing"
import "reflect"

func TestExtractNodes(t *testing.T){
	inString := "<a>Test</a>"
	outArray := []string{inString}
	extractedNodes := ExtractNodes(inString)
	if !reflect.DeepEqual(outArray,extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inString, extractedNodes, outArray)
	}

}
