package utils

import (
	"encoding/json"
	"log"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
