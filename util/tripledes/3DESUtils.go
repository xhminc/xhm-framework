package tripledes

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"errors"
	"github.com/xhminc/xhm-framework/config"
	"log"
	"os"
	"runtime"
)

const (
	ivdes = "xhminc20"
)

func EncryptToStringWithEnv(text string) string {
	str, err := EncryptToString(text, os.Getenv(config.DES_KEY), []byte(os.Getenv(config.DES_IV))...)
	if err != nil {
		panic(err)
	}
	return str
}

func EncryptToString(text string, key string, iv ...byte) (string, error) {
	encrypt, err := Encrypt([]byte(text), []byte(key), iv...)
	return hex.EncodeToString(encrypt), err
}

func Encrypt(text []byte, key []byte, ivDes ...byte) ([]byte, error) {

	if len(key) != 24 {
		return nil, errors.New("a twenty-four-length secret key is required")
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	paddingText := PKCS5Padding(text, block.BlockSize())

	var iv []byte
	if len(ivDes) != 0 {
		if len(ivDes) != 8 {
			return nil, errors.New("a eight-length ivdes key is required")
		} else {
			iv = ivDes
		}
	} else {
		iv = []byte(ivdes)
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)

	return cipherText, nil
}

func DecryptToString(text string, key string, iv ...byte) (string, error) {

	decodeString, decryptErr := hex.DecodeString(text)
	if decryptErr != nil {
		return "", decryptErr
	}

	decrypt, err := Decrypt(decodeString, []byte(key), iv...)
	return string(decrypt), err
}

func Decrypt(text []byte, key []byte, ivDes ...byte) ([]byte, error) {

	if len(key) != 24 {
		return nil, errors.New("a twenty-four-length secret key is required")
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime error: ", err, "Check that the key is correct")
			default:
				log.Println("error: ", err)
			}
		}
	}()

	var iv []byte
	if len(ivDes) != 0 {
		if len(ivDes) != 8 {
			return nil, errors.New("a eight-length ivdes key is required")
		} else {
			iv = ivDes
		}
	} else {
		iv = []byte(ivdes)
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, len(text))
	blockMode.CryptBlocks(paddingText, text)
	plainText, err := PKCS5UnPadding(paddingText)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}
