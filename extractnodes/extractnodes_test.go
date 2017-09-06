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
			ProgramOptions{TagToLookFor: EMPTY_STRING, FilterToApply: EMPTY_STRING, RetainTags: false, NameSpace: EMPTY_STRING},
			[]string{`<B>Test1</B>`, `<B>Test2</B>`},
		},
		{defaultInputString,
			`A`,
			ProgramOptions{TagToLookFor: EMPTY_STRING, FilterToApply: EMPTY_STRING, RetainTags: true, NameSpace: EMPTY_STRING},
			[]string{`<A><B>Test1</B></A>`, `<A><B>Test2</B></A>`},
		},
		{defaultInputString,
			`A`,
			ProgramOptions{TagToLookFor: `B`, FilterToApply: `Test1`, RetainTags: false, NameSpace: EMPTY_STRING},
			[]string{`<B>Test1</B>`},
		},
		{`<data><A xmlns="namespace1"><B>Test1</B></A><A xmlns="namespace2"><B>Test2</B></A></data>`,
			`A`,
			ProgramOptions{TagToLookFor: `B`, FilterToApply: EMPTY_STRING, RetainTags: false, NameSpace: "namespace1"},
			[]string{`<B>Test1</B>`},
		},
		{`<A xmlns="namespace1"><B>Test1</B></A>`,
			`A`,
			ProgramOptions{TagToLookFor: `B`, FilterToApply: EMPTY_STRING, RetainTags: true, NameSpace: "namespace1"},
			[]string{`<A xmlns="namespace1"><B>Test1</B></A>`},
		},
		{`<data><A xmlns="namespace1"><B>Test1</B></A></data>`,
			`B`,
			ProgramOptions{TagToLookFor: EMPTY_STRING, FilterToApply: EMPTY_STRING, RetainTags: false, NameSpace: "namespace1"},
			[]string{`Test1`},
		},
		{`<data xmlns:namespace1="http://namespace1.org"><namespace1:A><B>Test1</B></namespace1:A></data>`,
			`A`,
			ProgramOptions{TagToLookFor: EMPTY_STRING, FilterToApply: EMPTY_STRING, RetainTags: false, NameSpace: "namespace1"},
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

	programOptions := []ProgramOptions{
		{EMPTY_STRING, EMPTY_STRING, false, EMPTY_STRING},
		{`B`, `Test1`, false, EMPTY_STRING},
		{`B`, `Test*`, false, EMPTY_STRING},
		{`B`, EMPTY_STRING, false, EMPTY_STRING},
	}
	for _, options := range programOptions {
		weWantThisNode := WeWantThisNode(node, options)
		if weWantThisNode != true {
			t.Errorf("WeWantThisNode returned false for %s with options: %s", node, options)
		}
	}

}
