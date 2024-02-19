package main

import (
	"fmt"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"image/jpeg"
	"os"
)

func main() {
	inputFile := "input.jpg"
	outputFile := "output.webp"
	outputFile2 := "output2.webp"

	// 打开 JPEG 文件
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 读取 JPEG 图像
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建 WebP 文件
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer output.Close()

	// 编码图像并写入 WebP 文件
	err = webp.Encode(output, img, &webp.Options{Lossless: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建 WebP 文件
	output2, err2 := os.Create(outputFile2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer output2.Close()

	img2 := imaging.Resize(img, 320, 320, imaging.Lanczos)
	err = webp.Encode(output2, img2, &webp.Options{Lossless: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("转换成功!")
}
