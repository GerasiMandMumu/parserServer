package tags

import "encoding/xml"

type Model struct {
	XMLName     xml.Name   `xml:"model"`
	Types       Types      `xml:"types"`
	Description string     `xml:"description"`
	Author      string     `xml:"author"`
	Version     string     `xml:"version"`
	Imports     Imports    `xml:"imports"`
	Namespaces  Namespaces `xml:"namespaces"`
	Name        string     `xml:"name,attr"`
}
