package main

import (
	"log"
)


func main() {
	url := "https://catfact.ninja/fact"
 
	catFactService := NewCatFactService(url)
 catFactService = NewLoggingService(catFactService)
 
 apiServer := NewApiServer(catFactService)
 err := apiServer.Start(":8080")

 log.Fatal(err)

qrService := NewQRCodeService() 
err = qrService.GenerateQRCode(url, "cat-fact.png")
if err != nil {
    log.Fatal("Failed to generate QR code:", err)
}
log.Println("QR code generated!")
}
