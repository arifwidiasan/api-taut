package helper

import (
	"github.com/arifwidiasan/api-taut/config"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRCode(username string) (string, error) {
	fileName := username + ".png"
	err := qrcode.WriteFile(config.InitConfiguration().QR_PAGE+username, qrcode.Medium, 256, "../uploads/qrcode/"+fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
