package util

import (
	"log"
	"encoding/hex"
	"encoding/base64"
)

func HexToBase64(str string) string {
	decoded, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal("Error decoding string")
	}
	return base64.StdEncoding.EncodeToString(decoded)
}