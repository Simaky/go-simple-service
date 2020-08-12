package service

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"os"

	"github.com/nfnt/resize"
)

const (
	imagesPath = "./static/images/"
	maxWidth   = 160
	maxHeight  = 160

	fileFormat = ".jpg"
)

// ImageUpload save avatar to local storage
func ImageUpload(file multipart.File, userID uint) (string, error) {
	img, config, err := getImageAndConfig(file)
	if err != nil {
		return "", err
	}

	img, err = cropImage(img, config)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%d_user%s", userID, fileFormat)
	f, err := os.Create(imagesPath + fileName)
	if err != nil {
		return "", err
	}

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func cropImage(img image.Image, config image.Config) (image.Image, error) {
	if config.Height > maxHeight || config.Width > maxWidth {
		return resize.Resize(maxWidth, maxHeight, img, resize.Lanczos3), nil
	}
	return img, nil
}

// TODO add support more image formats (for now supports only .jpg)
func getImageAndConfig(file multipart.File) (image.Image, image.Config, error) {
	var secondCopy bytes.Buffer
	firstCopy := io.TeeReader(file, &secondCopy)

	img, err := jpeg.Decode(firstCopy)
	if err != nil {
		return img, image.Config{}, err
	}
	config, err := jpeg.DecodeConfig(&secondCopy)
	return img, config, err
}
