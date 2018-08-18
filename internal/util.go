package util

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBase64(str string) string {
	decoded, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal("Error decoding string")
	}
	return base64.StdEncoding.EncodeToString(decoded)
}

func HexToFixedXor(str string, against string) string {
	decoded, err := hex.DecodeString(str)
	decoded_against, err := hex.DecodeString(against)
	if err != nil {
		log.Fatal("Error decoding string")
	}
	var xord []byte
	for i := 0; i < len(decoded); i++ {
		xord = append(xord, decoded[i]^decoded_against[i])
	}
	return hex.EncodeToString(xord)
}
