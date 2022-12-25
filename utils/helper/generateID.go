package helper

import (
	"strings"
	"crypto/rand"
	"encoding/hex"
	)
	
func GenerateID() string  {
	id := make([]byte, 5)
	_, err := rand.Read(id)
	if err != nil {
		panic(err)
	}
	return "BOLAXD" + strings.ToUpper(hex.EncodeToString(id))
}

