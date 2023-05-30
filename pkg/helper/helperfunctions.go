package helper

import (
	"encoding/base64"
	"log"
	"math/big"
	"math/rand"
	"time"

	ra "crypto/rand"
	"fmt"
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

func GenerateTransactionID() string {
	randomID, _ := ra.Int(ra.Reader, big.NewInt(1e8-1)) // Generate a random number between 0 and 99999999 (8 digits)
	transactionID := fmt.Sprintf("%08d", randomID)

	return transactionID
}
