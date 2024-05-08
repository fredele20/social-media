package utils

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(hashedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}

func ConvertStruct(data interface{}, result interface{}) error {
	transferRateByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(transferRateByte, &result); err != nil {
		return err
	}

	return nil
}

func ConvertDataToMap(data interface{}) map[string]interface{} {
	var f interface{}
	ConvertStruct(&data, &f)

	return f.(map[string]interface{})
}

func GenerateId() string {
	return primitive.NewObjectID().Hex()
}
