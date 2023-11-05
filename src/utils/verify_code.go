package utils

import (
	"bytes"
	"embed"
	"encoding/base64"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
)

var (
	Height          = 50
	Width           = 100
	NoiseCount      = 4
	ShowLineOptions = base64Captcha.OptionShowHollowLine
	Length          = 4
	Source          = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	BgColor         = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	FontsStorage base64Captcha.FontsStorage
	Fonts        []string
)

type VerCode struct {
	ID          string
	Base64Image string
	Vcode       string
}

// 开辟空间（干什么用的不知道）
var store = base64Captcha.DefaultMemStore

// 查找包
//
//go:embed fontDirs/*
var embedFonts embed.FS

func (v *VerCode) CreatVerifyCode() {
	driver := base64Captcha.NewDriverString(Height, Width, Length, ShowLineOptions, NoiseCount, Source, &BgColor, FontsStorage, Fonts)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	v.ID, v.Base64Image, _ = captcha.Generate()
	v.Base64Image = v.Base64Image[22:]
	v.Vcode = captcha.Store.Get(v.ID, true)
}

func (v *VerCode) OutInput() []byte {
	base64String := v.Base64Image
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatal("Cannot decode b64")
	}
	reader := bytes.NewReader(decodedBytes)
	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal("Bad png")
	}
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, nil)
	if err != nil {
		log.Fatal("Cannot encode image")
	}
	return buf.Bytes()
}
