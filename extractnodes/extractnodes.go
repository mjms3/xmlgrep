package extractnodes

import (
	"encoding/xml"
	"io"
)

type InnerXmlContent struct {
	UnderlyingString string `xml:",innerxml"`
}

func ExtractNodes(inputXml io.Reader, targetTag string) []string  {
	var innerXml InnerXmlContent
	decoder := xml.NewDecoder(inputXml)
	nodesList := make([]string,0,0)
	for {
		token,_:= decoder.Token()
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
