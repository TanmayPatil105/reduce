package main

import (
	"github.com/skip2/go-qrcode"
)

func createQr(url string, file_name string) {

	if file_name == "" {
		file_name = "QR"
	}

	err := qrcode.WriteFile(url, qrcode.Medium, 156, file_name+".png")

	if err != nil {
		Box("QR-CODE", "Couldn't create QR code")
	}
}
