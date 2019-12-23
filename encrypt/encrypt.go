package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"final-project/constant"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Data(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher(CreateHash(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText, nil
}

func File(filename string, data []byte, passphrase string) ([]byte, error) {
	f, _ := os.Create(filename)
	defer f.Close()
	cipherText, err := Data(data, passphrase)
	if err != nil {
		return []byte{}, err
	}
	f.Write(cipherText)
	return cipherText, nil
}
func CreateHash(key string) []byte {
	salt := make([]byte, constant.PW_SALT_BYTES)
	return pbkdf2.Key([]byte(key), salt, 4096, 32, sha1.New)
}
