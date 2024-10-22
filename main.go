package main

import (
	"log"

	"github.com/pquerna/otp/totp"
)

func main() {
	// Generate a new secret
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "",
		AccountName: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output the URL for the QR code
	log.Printf("Key (base32): %s", key.Secret())
	log.Printf("QR Code URL: %s", key.URL())
}
