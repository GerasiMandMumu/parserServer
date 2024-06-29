package Utils

import (
	"fmt"
	"myproj/Config"
)

/*
* Функция чтения файла с конфигом
 */
func ReadPropertyFile() Config.AppConfigProperties {
	fileName := "global.properties"

	props, err := Config.ReadPropertiesFile(fileName)
	if err != nil {
		fmt.Println("Error read " + fileName)
	}
	return props
}
