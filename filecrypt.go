package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func encryptFile(key []byte, fileIn, fileOut string) error {
	plaintext, err := os.ReadFile(fileIn)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	head, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, head.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	cipher := head.Seal(nonce, nonce, plaintext, nil)

	return os.WriteFile(fileOut, cipher, 0644)
}

func decryptFile(key []byte, fileIn, fileOut string) error {
	ciphertext, err := os.ReadFile(fileIn)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := aead.NonceSize()

	if len(ciphertext) < nonceSize {
		return fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:nonceSize]
	cipherTextOnly := ciphertext[nonceSize:]

	plaintext, err := aead.Open(nil, nonce, cipherTextOnly, nil)
	if err != nil {
		return err
	}

	return os.WriteFile(fileOut, plaintext, 0644)
}
