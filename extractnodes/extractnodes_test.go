package extractnodes

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtractNodes(t *testing.T) {
	inputString := `<data><A><B>Test1</B></A><A><B>Test2</B></A></data>`
	inString := strings.NewReader(inputString)
	params := FilteringParams{TagToLookFor: EMPTY_STRING, FilterToApply: EMPTY_STRING}
	extractedNodes := ExtractNodes(inString, "A", params)
	expectedOutput := []string{`<B>Test1</B>`, `<B>Test2</B>`}
	if !reflect.DeepEqual(expectedOutput, extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inputString, extractedNodes, expectedOutput)
	}

}

func TestExtractNodes_withFilter(t *testing.T) {
	inputString := `<data><A><B>Test1</B></A><A><B>Test2</B></A></data>`
	inString := strings.NewReader(inputString)
	params := FilteringParams{TagToLookFor: "B", FilterToApply: "Test1"}
	extractedNodes := ExtractNodes(inString, "A", params)
	expectedOutput := []string{`<B>Test1</B>`}
	if !reflect.DeepEqual(expectedOutput, extractedNodes) {
		t.Errorf("ExtractNodes with input %s returned %s, expected: %s", inputString, extractedNodes, expectedOutput)
	}

}

func TestWeWantThisNode(t *testing.T) {
	node := "<B>Test1</B>"

	testInputs := []FilteringParams{
		{EMPTY_STRING, EMPTY_STRING},
		{"B", "Test1"},
		{"B", "Test*"},
		{"B", EMPTY_STRING},
	}
	for _, inputs := range testInputs {
		weWantThisNode := WeWantThisNode(node, inputs.TagToLookFor, inputs.FilterToApply)
		if weWantThisNode != true {
			t.Errorf("WeWantThisNode returned false for %s with tag: %s and regex: %s", node,
				inputs.TagToLookFor, inputs.FilterToApply)
		}
	}

}
