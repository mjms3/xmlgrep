package main

import (
	"bytes"
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
	fs := &testFs{
		fileContents: make(map[string]string),
	}

	fs.fileContents["file1"] = "<A>test1</A>"
	os.Args = []string{"DummyFileName", "A", "file1"}
	outputBuffer := new(bytes.Buffer)
	run(fs, outputBuffer)
	outputString := outputBuffer.String()
	expectedString := "test1\n"
	if outputString != expectedString {
		t.Errorf("Integration test with file contents: %s, args: %s, returned: %s, expected: %s",
			fs.fileContents, os.Args, outputString, expectedString)
	}
}
