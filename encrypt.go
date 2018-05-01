package main

import (
	"fmt"
	"io"
	"flag"
	"errors"

    "crypto/aes"
    "crypto/hmac"
	"crypto/rand"
	"crypto/cipher"
    "crypto/sha256"
    "crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

func encrypt(plaintext, key []byte) (string, error) {

	encryptionKey, salt, err := genKey(key)

    if err != nil {
        panic(err.Error())
    }
    block, err := aes.NewCipher(encryptionKey)
    if err != nil {
        panic(err.Error())
	}
	
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	
	nonce  := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	hmac := genHMAC(ciphertext, encryptionKey)

	ciphertext = append(hmac, ciphertext...)
	ciphertext = append(salt, ciphertext...)
	ciphertext = append(nonce, ciphertext...)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(b64ciphertext string, key []byte) ([]byte, error) {
    ciphertext, err := base64.StdEncoding.DecodeString(b64ciphertext)
    if err != nil {
        return ciphertext, err
    }
	nonce, salt, hmac, ciphertext := ciphertext[:12],ciphertext[12:76], ciphertext[76:108], ciphertext[108:]
	
	encryptionKey, err :=  genPBKDF2Key(key, salt)
	if err != nil {
		panic(err.Error())
	}

    hmacNew := genHMAC(ciphertext, encryptionKey)
    if subtle.ConstantTimeCompare(hmac, hmacNew) != 1 {
        return ciphertext, errors.New("invalid hmac")
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		panic(err.Error())
	}
	
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
    return plaintext, nil
}

func genHMAC(message, key []byte) []byte {
    mac := hmac.New(sha256.New, key)
    mac.Write(message)
    return mac.Sum(nil)
}

func genKey(key []byte) ([]byte, []byte,  error) {
    salt := make([]byte, 64)
    if _, err := rand.Read(salt); err != nil {
        return nil, nil, err
	}

	genkey, err := genPBKDF2Key(key, salt)
	if err != nil {
		panic(err.Error())
	}

	return genkey, salt, nil
}

func genPBKDF2Key(key []byte, salt []byte) ( []byte, error) {
	pbkdf2key := pbkdf2.Key(key, salt, 1000, 16, sha256.New)
	return pbkdf2key, nil
}

func main() {
	var key []byte
	msgPtr := flag.String("message", "This is my secret message", "Enter the message to encrypt.")
	keyPtr := flag.String("key", "", "")
	flag.Parse()

	if *keyPtr == "" {
		key = make([]byte, 64)    
		if _, err := rand.Read(key); err != nil {
			panic(err.Error())
		}
	} else {
		key = []byte(*keyPtr)
	}
	
	encoded_msg, _ := encrypt([]byte(*msgPtr), key)
	fmt.Printf("'%s' is encrypted as - %s\n", *msgPtr, encoded_msg)
	fmt.Printf("Key - %s\n", base64.StdEncoding.EncodeToString(key))	

	new_msg, _ := decrypt(encoded_msg, key)
	fmt.Printf("Decoded string - '%s'\n", new_msg)
}