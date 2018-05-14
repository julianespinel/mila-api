package files

import (
	"fmt"
	"io"
	"os"
	"time"
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

func GetBVCTemporalFileName() string {
	fileName := "bvc-stocks-%v.xls"
	return fmt.Sprintf(fileName, time.Now().Unix())
}
