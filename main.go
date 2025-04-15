package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
)

func main() {
	mode := flag.String("mode", "", "Mode: encrypt, decrypt, genkey")
	in := flag.String("in", "", "Input file")
	out := flag.String("out", "", "Output file")
	keyHex := flag.String("key", "", "Hex-encoded key for encryption/decryption")
	flag.Parse()

	switch *mode {
	case "genkey":
		key := make([]byte, 32) // AES-256
		if _, err := rand.Read(key); err != nil {
			log.Fatalf("Failed to generate key: %v", err)
		}
		fmt.Println("Your AES-256 key (hex):")
		fmt.Println(hex.EncodeToString(key))
	case "encrypt":
		if *keyHex == "" || *in == "" || *out == "" {
			log.Fatal("Missing required flags: -key, -in, -out")
		}
		key, err := hex.DecodeString(*keyHex)
		if err != nil {
			log.Fatalf("Invalid hex key: %v", err)
		}
		err = encryptFile(key, *in, *out)
		if err != nil {
			log.Fatalf("Encryption failed: %v", err)
		}
		fmt.Println("Encryption successful.")
	case "decrypt":
		if *keyHex == "" || *in == "" || *out == "" {
			log.Fatal("Missing required flags: -key, -in, -out")
		}
		key, err := hex.DecodeString(*keyHex)
		if err != nil {
			log.Fatalf("Invalid hex key: %v", err)
		}
		err = decryptFile(key, *in, *out)
		if err != nil {
			log.Fatalf("Decryption failed: %v", err)
		}
		fmt.Println("Decryption successful.")
	default:
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
}
