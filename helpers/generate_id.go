package helpers

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	uuidString := uuidObj.String()
	return uuidString, nil
}
