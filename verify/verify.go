package main

import (
	"log"

	"github.com/pquerna/otp/totp"
)

func verifyCode(secret, code string) bool {
	return totp.Validate(code, secret)
}

func main() {
	secret := ""
	code := "" // OTP Code

	if verifyCode(secret, code) {
		log.Println("Kode valid!")
	} else {
		log.Println("Kode tidak valid.")
	}
}
