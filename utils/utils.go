package utils

import (
	"encoding/json"
	"log"
)

func StructToJson(src any) string {
	jsonBytes, err := json.Marshal(src)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}

	return string(jsonBytes)
}

func JsonToSruct(src []byte, to any) bool {
	err := json.Unmarshal(src, &to)
	if err != nil {
		log.Println("Error:", err)
		return false
	}

	return true
}
