package authentic

import (
	"github.com/gofrs/uuid"
	"log"
)

func generateSessionId() string {
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u2)

	return u2.String()
}
