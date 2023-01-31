package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveFile(c *gin.Context, key string) (*string, error) {
	var filePath string
	file, header, err := c.Request.FormFile(key)
	if err != nil {
		return nil, err
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	baseDir := filepath.Join("data", key)
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	filePath = filepath.Join(baseDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return nil, err
	}
	return &filename, nil
}
