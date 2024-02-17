package main

import (
	"context"
	"fmt"
	"log"
)


func main() {
	url := "https://catfact.ninja/fact"
 catFactService := NewCatFactService(url)
 catFact, err := catFactService.GetCatFact(context.TODO())
 if err != nil {
	log.Fatal(err)
 }

 fmt.Println(catFact.Fact)

qrService := NewQRCodeService() 
err = qrService.GenerateQRCode(url, "cat-fact.png")
if err != nil {
    log.Fatal("Failed to generate QR code:", err)
}

}
