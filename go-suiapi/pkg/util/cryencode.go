package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/bwmarrin/snowflake"
)

var cryenkey = []byte{119, 144, 222, 171, 144, 248, 2, 157, 71, 22, 32, 44, 192, 144, 156, 213, 33, 198, 227, 233, 177, 213, 110, 210, 209, 141, 12, 194, 231, 62, 51, 72}

// 生成一个随机密钥
func generateKey() ([]byte, error) {
	key := make([]byte, 32) // 32字节对应AES-256
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func gennewid() (genid string) {
	n, err := snowflake.NewNode(2024)
	if err != nil {
		log.Fatal(2, " snowflake error : %v", err)
	}
	return n.Generate().String()
}

func engennewid() (genid string) {
	newid := gennewid()
	cipherText, err := Encrypt([]byte(newid))
	if err != nil {
		fmt.Println("Error encrypting text:", err)
		return
	}
	return cipherText
}

// AES加密
func Encrypt(plainText []byte) (string, error) {
	block, err := aes.NewCipher(cryenkey)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText), nil
}

// AES解密
func Decrypt(cipherTextHex string) (string, error) {
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(cryenkey)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

//func main() {
//	key, err := generateKey()
//	if err != nil {
//		fmt.Println("Error generating key:", err)
//		return
//	}
//	//[119 144 222 171 144 248 2 157 71 22 32 44 192 144 156 213 33 198 227 233 177 213 110 210 209 141 12 194 231 62 51 72]
//	{119,144,222,171,144,248,2,157,71,22,32,44,192,144,156,213,33,198,227,233,177,213,110,210,209,141,12,194,231,62,51,72}
//	plainText := "Hello, Golang encryption!"
//	fmt.Println("Original text:", plainText)
//
//	cipherText, err := encrypt([]byte(plainText), key)
//	if err != nil {
//		fmt.Println("Error encrypting text:", err)
//		return
//	}
//	fmt.Println("Encrypted text:", cipherText)
//
//	decryptedText, err := decrypt(cipherText, key)
//	if err != nil {
//		fmt.Println("Error decrypting text:", err)
//		return
//	}
//	fmt.Println("Decrypted text:", decryptedText)
//}
