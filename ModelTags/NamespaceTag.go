package tags

import "encoding/xml"

type Namespaces struct {
	XMLName    xml.Name    `xml:"namespaces"`
	Namespaces []Namespace `xml:"namespace"`
}

type Namespace struct {
	XMLName xml.Name `xml:"namespace"`
	Uri     string   `xml:"uri,attr"`
	Prefix  string   `xml:"prefix,attr"`
}
