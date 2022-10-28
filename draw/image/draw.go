package main

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
)

//DrawTextInfo 图片绘字信息
type DrawTextInfo struct {
	Text string
	X    int
	Y    int
}

//TextBrush 字体相关
type TextBrush struct {
	FontType  *truetype.Font
	FontSize  float64
	FontColor *image.Uniform
	TextWidth int
}

//NewTextBrush 新生成笔刷
func NewTextBrush(FontFilePath string, FontSize float64, FontColor *image.Uniform, textWidth int) (*TextBrush, error) {
	fontFile, err := ioutil.ReadFile(FontFilePath)
	if err != nil {
		return nil, err
	}
	fontType, err := truetype.Parse(fontFile)
	if err != nil {
		return nil, err
	}
	if textWidth <= 0 {
		textWidth = 20
	}
	return &TextBrush{FontType: fontType, FontSize: FontSize, FontColor: FontColor, TextWidth: textWidth}, nil
}

//Image2RGBA Image2RGBA
func Image2RGBA(img image.Image) *image.RGBA {

	baseSrcBounds := img.Bounds().Max

	newWidth := baseSrcBounds.X
	newHeight := baseSrcBounds.Y

	des := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight)) // 底板
	//首先将一个图片信息存入jpg
	draw.Draw(des, des.Bounds(), img, img.Bounds().Min, draw.Over)

	return des
}

//DrawStringOnImageAndSave 图片上写文字
func DrawStringOnImageAndSave(imageData []byte, infos []*DrawTextInfo) (*image.RGBA, error) {
	var err error
	//判断图片类型
	var backgroud image.Image
	filetype := http.DetectContentType(imageData)
	switch filetype {
	case "image/jpeg", "image/jpg":
		backgroud, err = jpeg.Decode(bytes.NewReader(imageData))
		if err != nil {
			fmt.Println("jpeg error")
			return nil, err
		}

	case "image/gif":
		backgroud, err = gif.Decode(bytes.NewReader(imageData))
		if err != nil {
			return nil, err
		}

	case "image/png":
		backgroud, err = png.Decode(bytes.NewReader(imageData))
		if err != nil {
			return nil, err
		}
	default:
		return nil, err
	}
	des := Image2RGBA(backgroud)

	//新建笔刷
	textBrush, _ := NewTextBrush("simfang.ttf", 20, image.Black, 50)

	//Px Py 绘图开始坐标 text要绘制的文字
	//调整颜色
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(textBrush.FontType)
	c.SetHinting(font.HintingFull)
	c.SetFontSize(textBrush.FontSize)
	c.SetClip(des.Bounds())
	c.SetDst(des)
	c.SetSrc(textBrush.FontColor)

	for _, info := range infos {
		c.DrawString(info.Text, freetype.Pt(info.X, info.Y))
	}

	return des, nil
}
