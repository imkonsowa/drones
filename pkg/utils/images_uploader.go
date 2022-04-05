package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"
)

const UploadsDirectoryPath = "storage/uploads"

// TODO: refactor to an uploading provider
func SaveImageFromBase64String(base64String string) (string, error) {
	ext, err := validateAndParse(base64String)
	if err != nil {
		return "", err
	}

	payload := strings.Index(base64String, ";base64,")
	dec, err := base64.StdEncoding.DecodeString(base64String[payload+8:])
	imageId := fmt.Sprintf("%s/%d_%d__%s", UploadsDirectoryPath, rand.Intn(100000000000), time.Now().Unix(), ext)

	r := bytes.NewReader(dec)

	switch ext {
	case ".png":
		image, err := png.Decode(r)
		if err != nil {
			return "", err
		}

		f, err := os.OpenFile(imageId, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		err = png.Encode(f, image)
		if err != nil {
			return "", err
		}
	case ".jpeg":
		image, err := jpeg.Decode(r)
		if err != nil {
			return "", err
		}

		f, err := os.OpenFile(imageId, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", err
		}

		err = jpeg.Encode(f, image, nil)
		if err != nil {
			return "", err
		}
	}

	return imageId, nil
}

func validateAndParse(base64String string) (string, error) {
	mimePrefix := strings.Split(base64String, ";")
	if len(mimePrefix) != 2 {
		return "", errors.New("can't inspect image mime type")
	}

	if strings.Contains(mimePrefix[0], "image/jpeg") {
		return ".jpeg", nil
	} else if strings.Contains(mimePrefix[0], "image/png") {
		return ".png", nil
	} else {
		return "", errors.New("not allowed image extension")
	}
}
