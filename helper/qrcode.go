package helper

import (
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRCode(username string) (string, error) {
	fileName := username + ".png"
	err := qrcode.WriteFile("https://www.localhost.com/profiles/"+username, qrcode.Medium, 256, "../uploads/qrcode/"+fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
