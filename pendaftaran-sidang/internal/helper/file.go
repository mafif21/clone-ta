package helper

import (
	"crypto/sha1"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

func FileHandler(err error, document *multipart.FileHeader) (string, error) {
	if err != nil {
		return "", err
	}

	fileExt := filepath.Ext(document.Filename)

	fileName := strings.TrimSuffix(document.Filename, fileExt)
	currentTime := time.Now().Format("20060102150405")
	fileNameWithTime := fmt.Sprintf("%s-%s", fileName, currentTime)

	sha := sha1.New()
	sha.Write([]byte(fileNameWithTime))
	encrypted := sha.Sum(nil)
	encryptedString := fmt.Sprintf("%x", encrypted)

	newFileName := fmt.Sprintf("%s%s", encryptedString, fileExt)
	return newFileName, nil
}
