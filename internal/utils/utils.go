package utils

import (
	"image"
	"io"
	"strings"

	_ "image/jpeg"
	_ "image/png"
)

func GetImageDimensions(file io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}

func IsImageFile(filename string) bool {
	ext := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])
	return ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "gif"
}
