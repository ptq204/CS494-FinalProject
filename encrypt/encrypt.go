package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// func Encrypt(key, text string) string {
// 	fmt.Println(text)
// 	block, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		panic(err)
// 	}
// 	plaintext := []byte(text)
// 	cfb := cipher.NewCFBEncrypter(block, iv)
// 	ciphertext := make([]byte, len(plaintext))
// 	cfb.XORKeyStream(ciphertext, plaintext)
// 	return encodeBase64(ciphertext)
// }
func Encrypt(key, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}
