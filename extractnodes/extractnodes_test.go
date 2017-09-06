package extractnodes

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtractNodes(t *testing.T) {
	type ExtractNodeTestCase struct {
		InputString    string
		TargetTag      string
		Options        ProgramOptions
		ExpectedOutput []string
	}

	defaultInputString := `<data><A><B>Test1</B></A><A><B>Test2</B></A></data>`

	testCases := []ExtractNodeTestCase{
		{defaultInputString,
			`A`,
			ProgramOptions{TagToLookFor: EMPTY_STRING, FilterToApply: EMPTY_STRING, RetainTags: false},
			[]string{`<B>Test1</B>`, `<B>Test2</B>`},
		},
		{defaultInputString,
			`A`,
			ProgramOptions{TagToLookFor: EMPTY_STRING, FilterToApply: EMPTY_STRING, RetainTags: true},
			[]string{`<A><B>Test1</B></A>`, `<A><B>Test2</B></A>`},
		},
		{defaultInputString,
			`A`,
			ProgramOptions{TagToLookFor: `B`, FilterToApply: `Test1`, RetainTags: false},
			[]string{`<B>Test1</B>`},
		},
	}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.InputString)
		extractedNodes := ExtractNodes(reader, testCase.TargetTag, testCase.Options)
		if !reflect.DeepEqual(testCase.ExpectedOutput, extractedNodes) {
			t.Errorf("ExtractedNodes(%s, %s, %s) returned: %s, expected: %s", testCase.InputString,
				testCase.TargetTag, testCase.Options, extractedNodes, testCase.ExpectedOutput)
		}
	}

}

func TestWeWantThisNode(t *testing.T) {
	node := `<B>Test1</B>`

	testInputs := []ProgramOptions{
		{EMPTY_STRING, EMPTY_STRING, false},
		{`B`, `Test1`, false},
		{`B`, `Test*`, false},
		{`B`, EMPTY_STRING, false},
	}
	for _, inputs := range testInputs {
		weWantThisNode := WeWantThisNode(node, inputs.TagToLookFor, inputs.FilterToApply)
		if weWantThisNode != true {
			t.Errorf("WeWantThisNode returned false for %s with inputs: %s", node, inputs)
		}
	}

}
