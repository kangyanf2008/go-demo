package main

import (
	"fmt"
	"github.com/3d0c/gmf"
)

func main() {
	ctx, err := gmf.NewInputCtx("input.mp4")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ctx.Free()

	// 这里你可以进行各种 FFmpeg 操作

	fmt.Println("FFmpeg context created successfully.")
}
