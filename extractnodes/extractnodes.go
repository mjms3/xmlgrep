package extractnodes

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"regexp"
	"strings"
	"github.com/renstrom/dedent"
)

const EMPTY_STRING string = ""

type InnerXmlContent struct {
	UnderlyingString string `xml:",innerxml"`
}

type CharData struct {
	Contents string `xml:",chardata"`
}

type ProgramOptions struct {
	TagToLookFor  string
	FilterToApply string
	RetainTags    bool
}

func ExtractNodes(inputXml io.Reader, targetTag string, params ProgramOptions) []string {
	var innerXml InnerXmlContent
	decoder := xml.NewDecoder(inputXml)
	nodesList := make([]string, 0, 0)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch currentElement := token.(type) {
		case xml.StartElement:
			if currentElement.Name.Local == targetTag {
				decoder.DecodeElement(&innerXml, &currentElement)
				if WeWantThisNode(innerXml.UnderlyingString, params.TagToLookFor,
					params.FilterToApply) == true {
					var buffer bytes.Buffer
					writer := bufio.NewWriter(&buffer)
					e := xml.NewEncoder(writer)
					if params.RetainTags {
						e.EncodeToken(currentElement)
					}
					writer.Write([]byte(dedent.Dedent(innerXml.UnderlyingString)))
					if params.RetainTags {
						e.EncodeToken(xml.EndElement{currentElement.Name})
					}
					e.Flush()
					writer.Flush()
					nodesList = append(nodesList, buffer.String())
				}

			}
		}
	}
	return nodesList
}

func WeWantThisNode(node string, tagOfInterest string, requiredValue string) bool {
	if tagOfInterest == EMPTY_STRING {
		return true
	}
	xmlAsReader := strings.NewReader(node)
	decoder := xml.NewDecoder(xmlAsReader)
	requiredValueRegex := regexp.MustCompile(requiredValue)
	var contents CharData
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch Element := token.(type) {
		case xml.StartElement:
			if Element.Name.Local == tagOfInterest {
				decoder.DecodeElement(&contents, &Element)
				if requiredValue == EMPTY_STRING || requiredValueRegex.MatchString(contents.Contents) {
					return true
				}
			}
		}
	}
	return false
}
