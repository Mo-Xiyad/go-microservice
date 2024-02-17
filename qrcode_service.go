package main

import (
	"github.com/skip2/go-qrcode"
)

type QRCodeGenerator interface {
    GenerateQRCode(data string, filename string) error
}


func NewQRCodeService() QRCodeGenerator {
    return &QRCodeService{}
}

type QRCodeService struct {}

func (s *QRCodeService) GenerateQRCode(data string, filename string) error {
    return qrcode.WriteFile(data, qrcode.Medium, 256, filename)
}
