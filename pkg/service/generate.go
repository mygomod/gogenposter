package service

import (
	"github.com/boombuler/barcode/qr"
	"os"
)

var url = "https://www.bagevent.com/event/gopherchina2020"

func Generate() (err error) {
	os.MkdirAll("./data/gen/thumb", os.ModePerm)
	os.MkdirAll("./data/gen/dst", os.ModePerm)
	os.MkdirAll("./data/gen/qrcode", os.ModePerm)

	// 生成二维码
	qrcode := NewQrCode(url, 140, 140, qr.M, qr.Auto)
	filePath, err := qrcode.Encode("./data/gen/qrcode")
	if err != nil {
		panic(err)
	}

	poster := NewPoster(
		Content{
			Title:   "Go的大数据",
			Author:  "毛剑",
			Company: "BILIBILI",
			BgPath:  "./data/img/poster.jpg",
			DstPath: "./data/gen/dst/poster.jpg",
		},
		&Rect{
			X0: 0,
			Y0: 0,
			X1: 750,
			Y1: 1334,
		},
		Avatar{
			Path:      "./data/img/maojian.jpeg",
			ThumbPath: "./data/gen/thumb/maojian.jpg",
			X:         59,
			Y:         192,
			Width:     632,
			Height:    627,
		},
		Qr{
			Path: filePath,
			X:    500,
			Y:    1058,
		},
	)

	err = poster.Generate()
	return
}
