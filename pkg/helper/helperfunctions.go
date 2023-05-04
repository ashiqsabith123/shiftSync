package helper

import (
	"encoding/base64"
	"log"
	"math/rand"
	"time"
)

func Decode(info string) []byte {
	data, err := base64.StdEncoding.DecodeString(info)

	if err != nil {
		log.Fatal("erorr when decoding")
	}

	return data
}

func CreateId() int {

	randomSource := rand.NewSource(time.Now().UnixNano())

	randomGenerator := rand.New(randomSource)

	randomInt := randomGenerator.Intn(100000)

	return randomInt
}
