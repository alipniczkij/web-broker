package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSON(fileName string) map[string][]string {
	datas := map[string][]string{}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &datas)

	return datas
}

func WriteJSON(fileName string, data map[string][]string) {
	jsonString, _ := json.Marshal(data)
	ioutil.WriteFile(fileName, jsonString, os.ModePerm)
}
