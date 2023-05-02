package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"shiftsync/pkg/config"
)

func createHash(key string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hasher.Sum(nil)
}

func Encrypt(data []byte) []byte {
	pass := config.GetCryptoSecret()

	block, _ := aes.NewCipher(createHash(pass))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte) []byte {
	pass := config.GetCryptoSecret()

	key := createHash(pass)
	block, err := aes.NewCipher(key)
	gcm, err1 := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err2 := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil || err1 != nil || err2 != nil {
		fmt.Println(err, err1, err2)
	}
	return plainText
}
