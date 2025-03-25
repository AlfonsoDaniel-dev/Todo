package uploads

import (
	"bytes"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func prepareImage(image image.Image, format string) ([]byte, error) {
	if format == "" {
		return nil, errors.New("Format is required")
	}

	var buf bytes.Buffer

	switch format {
	case "jpeg":
		err := jpeg.Encode(&buf, image, nil)
		if err != nil {
			return nil, err
		}
	case "png":
		err := png.Encode(&buf, image)
		if err != nil {
			return nil, err
		}

	}

	return buf.Bytes(), nil
}

func resizeImage(Data []byte, width uint) (image.Image, string, error) {
	if Data == nil || len(Data) == 0 || width <= 0 {
		return nil, "", errors.New("not enough arguments to resize image")
	}

	img, format, err := image.Decode(bytes.NewReader(Data))
	if err != nil {
		return nil, "", err
	}

	resizedImage := resize.Resize(width, 0, img, resize.Lanczos3)

	return resizedImage, format, nil
}

func uploadImage(ImageToUpload []byte, fileName, repositoryPath string) error {
	if ImageToUpload == nil || len(ImageToUpload) == 0 || fileName == "" {
		return errors.New("not data to upload image")
	}

	imgdecodedAndResized, format, err := resizeImage(ImageToUpload, 800)
	if err != nil {
		return err
	}

	img, err := prepareImage(imgdecodedAndResized, format)
	if err != nil {
		return err
	}

	file, err := os.Create(repositoryPath + "/" + fileName)
	if err != nil {
		return err
	}

	_, err = file.Write(img)
	if err != nil {
		return err
	}

	return nil
}

func getImage(fileName, repositoryPath string) ([]byte, error) {

	completeFile := fileName + ".jpeg"

	completePath := filepath.Join(repositoryPath, completeFile)

	return os.ReadFile(completePath)
}
