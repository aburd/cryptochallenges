package challenges

import (
	"encoding/hex"
	"log"
	"fmt"
	"github.com/aburd/cryptochallenge/internal"
)



func Init3() {
	var original string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decoded, err := hex.DecodeString(original)
	if err != nil {
		log.Fatal("Decoding was unsuccessful")
	}
	var characters []byte
	for i := 0; i <= 255; i++ {
		characters = append(characters, byte(i))
	}
	var results [][]byte
	for _, char := range characters {
		temp := make([]byte, len(characters))
		copy(temp, util.XorBytes(decoded, char))
		results = append(results, temp)
	}	
	var scored []float32
	for _, word := range results {
		scored = append(scored, util.ScoreWord(word))
	}
	maxScoreIndex := util.MaxIndex(scored)
	fmt.Println(string(results[maxScoreIndex]))
}
