package utils

import "github.com/google/uuid"

func GenerateUuid() string {
	randomUuid := uuid.New()
	return randomUuid.String()
}
