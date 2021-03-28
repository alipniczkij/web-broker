package tools

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSON(fileName string) (map[string][]string, error) {
	datas := map[string][]string{}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func WriteJSON(fileName string, data map[string][]string) error {
	jsonString, _ := json.Marshal(data)
	err := ioutil.WriteFile(fileName, jsonString, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
