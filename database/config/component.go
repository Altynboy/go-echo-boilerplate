package component

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigModel struct {
	DB struct {
		HOST   string `json:"HOST"`
		USERNAME   string `json:"USERNAME"`
		DATABASE string `json:"DATABASE"`
		PASSWORD string `json:"PASSWORD"`
		PORT   string    `json:"PORT"`
	} `json:"DB"`
	JWT struct {
		SECRET_KEY string `json:"SECRET_KEY"`
	} `json:"JWT"`
}

var ConfigData ConfigModel

func Config() *ConfigModel {
	return &ConfigModel{}
}


func (*ConfigModel) Init(fileName string) error {

	configFile, err := os.Open(fileName)
	defer configFile.Close()

	if err != nil {
		return fmt.Errorf("ConfigInit error: %+v", err)
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&ConfigData)
	if err != nil {
		return fmt.Errorf("Error parsing congfig file: %+v", err)
	}

	return nil
}

