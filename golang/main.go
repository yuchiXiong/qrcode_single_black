package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
	"github.com/chai2010/tiff"
)

func main() {
	// 生成二维码并获得字节数组
	qrCode, err := qrcode.Encode("https://intelligent-book-yuchi-e.suanshubang.cc/ibserver/video.html?nodeId=719471", qrcode.Medium, 1024)
	if err != nil {
		panic(err)
	}

	// 将字节数组转换为图像对象
	img, err := png.Decode(bytes.NewReader(qrCode))
	if err != nil {
		panic(err)
	}

	// 将图像转换为灰度模式
	grayImg := convertToGrayScale(img)

	// 保存灰度图像为 TIFF 文件
	outputPath := "output_image.tif"
	err = saveAsTiff(grayImg, outputPath)
	if err != nil {
		panic(err)
	}

	println("Grayscale TIFF image saved as:", outputPath)
}

// convertToGrayScale 将图像转换为灰度模式
func convertToGrayScale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	return grayImg
}

// saveAsTiff 保存图像为 TIFF 文件
func saveAsTiff(img image.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tiff.Encode(file, img, nil)
}
