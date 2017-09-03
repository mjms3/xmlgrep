package extractnodes

import (
	"encoding/xml"
	"fmt"
)

func ExtractNodes(inputXmlString string) []string  {
	type InnerXml struct {
		Content string `xml:",innerxml"`
	}

	type NodeList struct {
		Tag    []InnerXml     `xml:"A"`
	}

	result := NodeList{}
	err := xml.Unmarshal([]byte(inputXmlString), &result)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	innerXmlList := make([]string,len(result.Tag))
	for idx := range result.Tag {
		innerXmlList[idx] = result.Tag[idx].Content
	}

	return innerXmlList
}
