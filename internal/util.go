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


var EnglishLetterFrequencies = map[string]float32{
	"a": 8.167, "b": 1.492, "c": 2.782, "d": 4.253, "e": 12.702,
	"f": 2.228, "g": 2.015, "h": 6.094, "i": 6.966, "j": 0.153,
	"k": 0.772, "l": 4.025, "m": 2.406, "n": 6.749, "o": 7.507,
	"p": 1.929, "q": 0.095, "r": 5.987, "s": 6.327, "t": 9.056,
	"u": 2.758, "v": 2.360, "x": 0.150, "y": 1.974, "z": 0.074,
	" ": 13.0, // space is slightly more frequent than (e)
}

func ScoreWord(characters []byte) float32 {
	var wordScore float32
	for _, char := range characters {
		letter := string(char)
		wordScore += EnglishLetterFrequencies[letter]
	}
	return wordScore
}

func XorBytes(a []byte, k byte) []byte {
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ k
	}
	return res
}

func MaxIndex(v []float32) int {
	var (
		max int 		= 0
		currentHigh float32 = 0.0
	)
	if len(v) > 0 {
		max = 0
		currentHigh = v[0]
	}
	for i, e := range v {
		if e > currentHigh {
			max = i
			currentHigh =  e
		}
	}
    return max
}
