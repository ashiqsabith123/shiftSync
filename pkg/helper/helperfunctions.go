package helper

import (
	"encoding/base64"
	"log"
)

func Decode(info string) []byte {
	data, err := base64.StdEncoding.DecodeString(info)

	if err != nil {
		log.Fatal("erorr when decoding")
	}

	return data
}
