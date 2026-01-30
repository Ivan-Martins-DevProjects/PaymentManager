package security

import "github.com/google/uuid"

func GenerateToken() uuid.UUID {
	id := uuid.New()

	return id
}
