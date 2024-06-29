package tags

import "encoding/xml"

type MandatoryAspects struct {
	XMLName xml.Name `xml:"mandatory-aspects"`
	Aspects []Aspect `xml:"aspect"`
}

type Aspect struct {
	XMLName xml.Name `xml:"aspect"`
}
