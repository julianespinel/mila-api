package files

import (
	"io"
	"os"
)

func SaveBodyToFile(filePath string, body io.ReadCloser) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, body)
	return err
}
