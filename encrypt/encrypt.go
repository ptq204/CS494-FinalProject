package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(key, text string) string {
	fmt.Println(text)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext)
}
