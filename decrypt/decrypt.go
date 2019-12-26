package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func createHash(key string) []byte {
	hash := md5.Sum([]byte(key))
	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash[:])
	return dst
}
func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// func Decrypt(key, text string) string {
// 	fmt.Println(text)
// 	block, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		panic(err)
// 	}
// 	ciphertext := decodeBase64(text)
// 	cfb := cipher.NewCFBEncrypter(block, iv)
// 	plaintext := make([]byte, len(ciphertext))
// 	cfb.XORKeyStream(plaintext, ciphertext)
// 	return string(plaintext)
// }
func Decrypt(key string, text string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(text)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}
