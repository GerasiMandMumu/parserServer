package tags

import "encoding/xml"

type Imports struct {
	XMLName xml.Name `xml:"imports"`
	Imports []Import `xml:"import"`
}

type Import struct {
	XMLName xml.Name `xml:"import"`
	Uri     string   `xml:"uri,attr"`
	Prefix  string   `xml:"prefix,attr"`
}
