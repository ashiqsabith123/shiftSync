package helper

import (
	"encoding/base64"
	"fmt"
	"log"
)

func Decode(info string) []byte {
	data, err := base64.StdEncoding.DecodeString(info)

	if err != nil {
		log.Fatal("erorr when decoding")
	}

	fmt.Println(data)

	return data
}
