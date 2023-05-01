package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"shiftsync/pkg/config"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data []byte) []byte {
	pass := config.GetCryptoSecret()

	block, _ := aes.NewCipher([]byte(createHash(pass)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte) []byte {
	pass := config.GetCryptoSecret()

	key := []byte(createHash(pass))
	block, err := aes.NewCipher(key)
	gcm, err1 := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err2 := gcm.Open(nil, nonce, cipherText, nil)
	fmt.Println("1st", err)
	fmt.Println("2nd", err1)
	fmt.Println("3rd", err2)

	return plainText

}
