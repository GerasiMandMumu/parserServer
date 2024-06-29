package tags

import (
	"encoding/xml"
	"myproj/Constants/Other"
)

type Types struct {
	XMLName xml.Name `xml:"types"`
	Types   []Type   `xml:"type"`
}

type Type struct {
	XMLName          xml.Name         `xml:"type"`
	Title            string           `xml:"title"`
	Parent           string           `xml:"parent"`
	Properties       Properties       `xml:"properties"`
	MandatoryAspects MandatoryAspects `xml:"mandatory-aspects"`
	Name             string           `xml:"name,attr"`
}

// Получение массива названий свойств из типа
func (object *Type) GetPropertiesArray() []string {
	var properties []string
	properties = append(properties, Other.Name)
	for i := 0; i < len(object.Properties.Properties); i++ {
		var property = object.Properties.Properties[i]
		properties = append(properties, property.Name)
	}
	return properties
}
