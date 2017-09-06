package extractnodes

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"github.com/renstrom/dedent"
	"io"
	"reflect"
	"regexp"
	"strings"
	"unsafe"
)

const EMPTY_STRING string = ``

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
	NameSpace     string
}

func elementIsInNamespace(elt xml.StartElement, requiredNamespace string) bool {
	if requiredNamespace == EMPTY_STRING {
		return true
	}
	if elt.Name.Space == requiredNamespace {
		return true
	}
	return false
}

func ExtractNodes(inputXml io.Reader, targetTag string, params ProgramOptions) []string {
	var innerXml InnerXmlContent
	decoder := xml.NewDecoder(inputXml)
	unexportedNameSpaceMap := reflect.ValueOf(decoder).Elem().FieldByName("ns")
	pointerToUnexportedNameSpaceMap := reflect.NewAt(unexportedNameSpaceMap.Type(), unsafe.Pointer(unexportedNameSpaceMap.UnsafeAddr())).Elem()
	nameSpaceMap := pointerToUnexportedNameSpaceMap.Interface().(map[string]string)
	nodesList := make([]string, 0, 0)

	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch currentElement := token.(type) {
		case xml.StartElement:
			nameSpaceIdentifyingString, ok := nameSpaceMap[params.NameSpace]
			if !ok {
				nameSpaceIdentifyingString = params.NameSpace
			}
			if currentElement.Name.Local == targetTag && elementIsInNamespace(currentElement, nameSpaceIdentifyingString) {
				decoder.DecodeElement(&innerXml, &currentElement)
				if WeWantThisNode(innerXml.UnderlyingString, params) == true {
					var buffer bytes.Buffer
					writer := bufio.NewWriter(&buffer)
					e := xml.NewEncoder(writer)
					if params.RetainTags {
						// Deal with known golang bug 7535
						currentElement.Name.Space = ""
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

func WeWantThisNode(node string, options ProgramOptions) bool {
	if options.TagToLookFor == EMPTY_STRING {
		return true
	}
	xmlAsReader := strings.NewReader(node)
	decoder := xml.NewDecoder(xmlAsReader)
	requiredValueRegex := regexp.MustCompile(options.FilterToApply)
	var contents CharData
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch Element := token.(type) {
		case xml.StartElement:
			if Element.Name.Local == options.TagToLookFor {
				decoder.DecodeElement(&contents, &Element)
				if options.FilterToApply == EMPTY_STRING || requiredValueRegex.MatchString(contents.Contents) {
					return true
				}
			}
		}
	}
	return false
}
