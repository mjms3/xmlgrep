package extractnodes

import (
	"encoding/xml"
	"io"
	"regexp"
	"strings"
)

const EMPTY_STRING string = ""

type InnerXmlContent struct {
	UnderlyingString string `xml:",innerxml"`
}

type CharData struct {
	Contents string `xml:",chardata"`
}

type FilteringParams struct {
	TagToLookFor  string
	FilterToApply string
}

func ExtractNodes(inputXml io.Reader, targetTag string, params FilteringParams) []string {
	var innerXml InnerXmlContent
	decoder := xml.NewDecoder(inputXml)
	nodesList := make([]string, 0, 0)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch Element := token.(type) {
		case xml.StartElement:
			if Element.Name.Local == targetTag {
				decoder.DecodeElement(&innerXml, &Element)
				if WeWantThisNode(innerXml.UnderlyingString, params.TagToLookFor,
					params.FilterToApply) == true {
					nodesList = append(nodesList, innerXml.UnderlyingString)
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
