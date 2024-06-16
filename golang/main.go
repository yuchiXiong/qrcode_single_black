package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
	"github.com/sunshineplan/tiff"
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
	grayImg := ConvertToCMYK(img)

	// 保存灰度图像为 TIFF 文件
	outputPath := "origin_single_black.tif"
	err = SaveAsTiff(grayImg, outputPath)
	if err != nil {
		panic(err)
	}

	println("Grayscale TIFF image saved as:", outputPath)

	// // ----------------------------------------------------------------解决带彩色logo的问题----------------------------------------------------------------

	// var outputPath = "cmyk_0_0_0_100.tif"
	// 打开图像文件
	origin_tif_file, err := os.Open(outputPath)

	if err != nil {
		panic(err)
	}

	defer origin_tif_file.Close()
	// 解码 TIFF 文件
	origin_tif, err := tiff.Decode(origin_tif_file)

	if err != nil {
		panic(err)
	}

	// 打开水印文件
	wmb_file, err := os.Open("../watermark.png")

	if err != nil {
		panic(err)
	}

	defer wmb_file.Close()

	// 解码 PNG 文件
	wmb_img, err := png.Decode(wmb_file)

	if err != nil {
		panic(err)
	}

	//把水印写在正中间
	offset := image.Pt((origin_tif.Bounds().Dx()-wmb_img.Bounds().Dx())/2, (origin_tif.Bounds().Dy()-wmb_img.Bounds().Dy())/2)
	b := origin_tif.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewCMYK(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖 在目标图像上（类似图层）
	draw.Draw(m, b, origin_tif, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)

	//生成新的二维码
	var newFileName = "single_black_with_logo_color.tif"

	// exportType => 2 彩色，转成 tif 格式导出
	// if exportType == 2 {
	err = SaveAsTiff(m, newFileName)
	// } else {
	// 	// exportType => 1 黑白 转成灰度模式 tif 格式导出
	// 	grayImg := ConvertToGrayScale(m)
	// 	err = SaveAsTiff(grayImg, newFileName)
	// }

}

// 将图像转换位 CMYK 0 0 0 1 模式
func ConvertToCMYK(img image.Image) *image.CMYK {
	bounds := img.Bounds()
	cmykImg := image.NewCMYK(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			cmykColor := color.CMYKModel.Convert(originalColor)
			cmykImg.Set(x, y, cmykColor)
		}
	}

	return cmykImg
}

// // convertToGrayScale 将图像转换为灰度模式
// func ConvertToGrayScale(img image.Image) *image.Gray {
// 	bounds := img.Bounds()
// 	grayImg := image.NewGray(bounds)

// 	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
// 		for x := bounds.Min.X; x < bounds.Max.X; x++ {
// 			originalColor := img.At(x, y)
// 			grayColor := color.GrayModel.Convert(originalColor)
// 			grayImg.Set(x, y, grayColor)
// 		}
// 	}

// 	return grayImg
// }

// saveAsTiff 保存图像为 TIFF 文件
func SaveAsTiff(img image.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// return tiff.Encode(file, img, nil)
	return tiff.Encode(file, img, &tiff.Options{Compression: tiff.LZW})
}
