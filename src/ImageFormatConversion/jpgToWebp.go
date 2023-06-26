package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/chai2010/webp"
)

func main() {
	inputFile := "input.jpg"
	outputFile := "output.webp"

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

	fmt.Println("转换成功!")
}
