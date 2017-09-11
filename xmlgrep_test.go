package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

type testFs struct {
	fileContents map[string]string
}

func (fs testFs) Open(name string) (io.Reader, error) {
	fileContents := fs.fileContents[name]
	return strings.NewReader(fileContents), nil
}

func TestXmlGrep(t *testing.T) {
	type testXmlGrepTestCase struct {
		fileContents   map[string]string
		args           []string
		expectedOutput string
		osStdin        string
	}

	testCases := []testXmlGrepTestCase{
		{
			fileContents:   map[string]string{"file1": "<A>test1</A>"},
			args:           []string{"A", "file1"},
			osStdin:        "",
			expectedOutput: "test1\n",
		},
		{
			fileContents:   map[string]string{"file1": "<A>test1</A>"},
			args:           []string{"-r", "A", "file1"},
			osStdin:        "",
			expectedOutput: "<A>test1</A>\n",
		},
		{
			fileContents: map[string]string{
				"file1": "<A>test1</A>",
				"file2": "<A>test2</A>",
			},
			args:           []string{"-r", "A", "file1", "file2"},
			osStdin:        "",
			expectedOutput: "<A>test1</A>\n<A>test2</A>\n",
		},
	}

	for _, testCase := range testCases {
		fs := &testFs{
			fileContents: testCase.fileContents,
		}
		os.Args = append([]string{"DummyFileName"}, testCase.args...)
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
		inputBuffer := strings.NewReader(testCase.osStdin)

		outputBuffer := new(bytes.Buffer)
		run(fs, outputBuffer, inputBuffer)
		outputString := outputBuffer.String()
		if outputString != testCase.expectedOutput {
			t.Errorf("Integration test with file contents: %s, args: %s, stdin: %s, returned: %s, expected: %s",
				fs.fileContents, os.Args, testCase.osStdin, outputString, testCase.expectedOutput)
		}
	}

}
