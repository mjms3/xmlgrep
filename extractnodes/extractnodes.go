package extractnodes

import (
	"encoding/xml"
	"io"
	"regexp"
	"strings"
)

type InnerXmlContent struct {
	UnderlyingString string `xml:",innerxml"`
}

type CharData struct {
	Contents string `xml:",chardata"`
}

func ExtractNodes(inputXml io.Reader, targetTag string) []string {
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
				nodesList = append(nodesList, innerXml.UnderlyingString)
			}
		}
	}
	return nodesList
}

func DoWeWantThisNode(node string, tagOfInterest string, requiredValue string) bool {
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
				if requiredValueRegex.MatchString(contents.Contents) {
					return true
				}
			}
		}
	}
	return false
}
