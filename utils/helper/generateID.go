package helper

import (
	"strings"
	"crypto/rand"
	"encoding/hex"
	)
	
func GenerateID() string  {
	id := make([]byte, 14)
	_, err := rand.Read(id)
	if err != nil {
		panic(err)
	}
	return strings.ToUpper(hex.EncodeToString(id)) + "BALL"
}

