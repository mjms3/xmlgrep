package extractnodes

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtractNodes(t *testing.T) {
	inputString := `<data><A><B>Test1</B></A><A><B>Test2</B></A></data>`
	inString := strings.NewReader(inputString)
	params := FilteringParams{TagToLookFor: "",
		FilterToApply: ""}
	extractedNodes := ExtractNodes(inString, "A", params)
	expectedOutput := []string{`<B>Test1</B>`, `<B>Test2</B>`}
	if !reflect.DeepEqual(expectedOutput, extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inputString, extractedNodes, expectedOutput)
	}

}

func TestExtractNodes_withFilter(t *testing.T) {
	inputString := `<data><A><B>Test1</B></A><A><B>Test2</B></A></data>`
	inString := strings.NewReader(inputString)
	params := FilteringParams{TagToLookFor: "B",
		FilterToApply: "Test1"}
	extractedNodes := ExtractNodes(inString, "A", params)
	expectedOutput := []string{`<B>Test1</B>`}
	if !reflect.DeepEqual(expectedOutput, extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inputString, extractedNodes, expectedOutput)
	}

}

func TestDoWeWantThisNode_withNoTag_returnsTrue(t *testing.T) {
	node := "<B>Test1</B>"
	weWantThisNode := DoWeWantThisNode(node, "", "")
	if weWantThisNode != true {
		t.Errorf("DoWeWantThisNode returned false for %s", node)
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

func TestDoWeWantThisNode_withNoFilterOnlyLooksForTag(t *testing.T) {
	node := "<B>Test1</B>"
	weWantThisNode := DoWeWantThisNode(node, "B", "")
	if weWantThisNode != true {
		t.Errorf("DoWeWantThisNode returned false for %s", node)
	}
}
