package challenges

import (
	"encoding/hex"
	"bytes"
	"log"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/aburd/cryptochallenge/internal"
)

func findBestMatch(str string) []byte {
	decoded, err := hex.DecodeString(str)
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
	return findBestScore(results)
}

func findBestScore(wordList [][]byte) []byte {
	var scored []float32
	for _, word := range wordList {
		scored = append(scored, util.ScoreWord(word))
	}
	maxScoreIndex := util.MaxIndex(scored)
	return wordList[maxScoreIndex]
}

func Init4() {
	// Get file contents
	path, _ := filepath.Abs("level1/onecharencrypt.txt")
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("There was an err reading from the file.")
	}

	// Format contents
	lines := bytes.Split(dat, []byte("\n"))

	var bestMatches [][]byte
	for _, line := range lines {
		temp := make([]byte, len(line))
		copy(temp, findBestMatch(string(line)))
		bestMatches = append(bestMatches, temp)
	}
	final := findBestScore(bestMatches)
	fmt.Println(string(final))
}
