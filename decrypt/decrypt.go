package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"final-project/constant"
	"io/ioutil"

	"golang.org/x/crypto/pbkdf2"
)

func Data(data []byte, passphrase string) ([]byte, error) {
	key := CreateHash(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return []byte{}, err
	}
	return plainText, nil
}

func File(filename string, passphrase string) ([]byte, error) {
	data, _ := ioutil.ReadFile(filename)
	plainText, err := Data(data, passphrase)
	if err != nil {
		return []byte{}, err
	}
	return plainText, nil
}
func CreateHash(key string) []byte {
	salt := make([]byte, constant.PW_SALT_BYTES)
	return pbkdf2.Key([]byte(key), salt, 4096, 32, sha1.New)
}
