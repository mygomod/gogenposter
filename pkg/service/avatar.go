package service

import (
	"github.com/disintegration/imaging"
	"github.com/mygomod/gogenposter/pkg/mus"
	"go.uber.org/zap"
	"image/jpeg"
	"os"
)

type Avatar struct {
	Path      string
	ThumbPath string
	ThumbFile *os.File
	X         int
	Y         int
	Width     int
	Height    int
}

func (a *Avatar) Thumb() (err error) {
	img, err := imaging.Open(a.Path)
	if err != nil {
		mus.Logger.Error("img open error", zap.String("err", err.Error()))
		return
	}

	thumb := imaging.Fill(img, a.Width, a.Height, imaging.Top, imaging.Lanczos)
	//thumb := imaging.Thumbnail(img, 100, 100, imaging.CatmullRom)
	//dst := imaging.New(a.Width, a.Height, color.NRGBA{0, 0, 0, 0})
	//dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
	a.ThumbFile, err = os.OpenFile(a.ThumbPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	err = jpeg.Encode(a.ThumbFile, thumb, nil)
	defer a.ThumbFile.Close()
	return
}
