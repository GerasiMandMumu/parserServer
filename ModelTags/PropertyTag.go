package tags

import "encoding/xml"

type Properties struct {
	XMLName    xml.Name   `xml:"properties"`
	Properties []Property `xml:"property"`
}

type Property struct {
	XMLName     xml.Name `xml:"property"`
	Description string   `xml:"description"`
	Typ         string   `xml:"type"`
	Name        string   `xml:"name,attr"`
}
