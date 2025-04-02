package main

import (
	downloader "ffmpeg-downloader"
	"fmt"
)

func main() {
	path, installed := downloader.IsStaticallyInstalled("ffmpeg-downloader")
	fmt.Println(path, installed)

	path, err := downloader.Download("ffmpeg-downloader")
	fmt.Println(path, err)
}
