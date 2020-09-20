package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/jpeg"
	"os"
	"path"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// Encode generate QR code
func (q *QrCode) Encode(folderPath string) (filePath string, err error) {
	name := encodeMD5(q.URL) + q.GetQrCodeExt()
	filePath = path.Join(folderPath, name)
	code, err := qr.Encode(q.URL, q.Level, q.Mode)
	if err != nil {
		return "", err
	}

	code, err = barcode.Scale(code, q.Width, q.Height)
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = jpeg.Encode(f, code, nil)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
