package image

import (
	"github.com/golang/freetype"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type Service struct {
}

func NewImageService() *Service {
	return &Service{}
}

func (i *Service) DrawText(inputFileName, topText, bottomText string) (string, error) {
	imgFile, err := os.Open(inputFileName)
	if err != nil {
		return "Не могу открыть изображение", err
	}
	defer imgFile.Close()

	fontBytes, err := os.ReadFile("/Users/almaz/GolandProjects/TestProjects/memogen/Roboto-Regular.ttf")
	if err != nil {
		return "Не могу открыть шрифт", err
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return "Не могу спарсить шрифт", err
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return "Не могу декордировать изображение", err
	}
	// создаем новый рисунок RGBA
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	// Создаем контекс freetype для рисования текста
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(30)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.NewUniform(color.White))

	pt := freetype.Pt(10, 40)
	_, err = c.DrawString(topText, pt)
	if err != nil {
		return "", err
	}

	pt = freetype.Pt(10, rgba.Bounds().Dy()-10)
	_, err = c.DrawString(bottomText, pt)
	if err != nil {
		return "", err
	}

	outFile, err := os.CreateTemp("", "meme")
	if err != nil {
		return "Не могу создать временный файл", err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, rgba, nil)
	if err != nil {
		return "Не могу сохранить рисунок", err
	}

	return outFile.Name(), nil
}
