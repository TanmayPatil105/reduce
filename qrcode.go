package main

import (
	"github.com/skip2/go-qrcode"
)

func createQr(url string){
	err := qrcode.WriteFile(url, qrcode.Medium, 156, "qr.png")

	if err!= nil{
		Box("QR-CODE","Couldn't create QR code")
	}
}
