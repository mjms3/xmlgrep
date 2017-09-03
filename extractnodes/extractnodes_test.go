package extractnodes

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtractNodes(t *testing.T) {
	inputString := `<data><A><B>Test1</B></A><A><B>Test2</B></A></data>`
	inString := strings.NewReader(inputString)
	expectedOutput := []string{`<B>Test1</B>`, `<B>Test2</B>`}
	extractedNodes := ExtractNodes(inString, "A")
	if !reflect.DeepEqual(expectedOutput, extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inputString, extractedNodes, expectedOutput)
	}

}

func TestDoWeWantThisNode(t *testing.T) {
	node := "<B>Test1</B>"
	weWantThisNode := DoWeWantThisNode(node, "B", "Test1")
	if weWantThisNode != true {
		t.Errorf("DoWeWantThisNode returned false for %s", node)
	}
}

func TestDoWeWantThisNode_pattern(t *testing.T) {
	node := "<B>Test1</B>"
	weWantThisNode := DoWeWantThisNode(node, "B", "Test*")
	if weWantThisNode != true {
		t.Errorf("DoWeWantThisNode returned false for %s", node)
	}
}
