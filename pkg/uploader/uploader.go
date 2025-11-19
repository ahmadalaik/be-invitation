package uploader

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UploadConfig struct {
	Destination string
	MAXSizeMB   int64
	AllowedExt  []string
}

type UploadResult struct {
	FileName string
	FilePath string
}

type Uploader struct {
	cfg UploadConfig
}

func NewUploader(cfg UploadConfig) *Uploader {
	os.MkdirAll(cfg.Destination, 0755)
	return &Uploader{cfg}
}

func (u *Uploader) UploadSingle(c *fiber.Ctx, file *multipart.FileHeader) (*UploadResult, error) {
	if err := u.ValidationFile(file); err != nil {
		return nil, err
	}

	ext := filepath.Ext(file.Filename)
	newName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join(u.cfg.Destination, newName)

	if err := c.SaveFile(file, savePath); err != nil {
		return nil, err
	}

	return &UploadResult{
		FileName: newName,
		FilePath: savePath,
	}, nil
}

func (u *Uploader) UploadMultiple(c *fiber.Ctx, files []*multipart.FileHeader) ([]UploadResult, error) {
	var results []UploadResult

	for _, file := range files {
		if err := u.ValidationFile(file); err != nil {
			return nil, err
		}

		ext := filepath.Ext(file.Filename)
		newName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		savePath := filepath.Join(u.cfg.Destination, newName)

		if err := c.SaveFile(file, savePath); err != nil {
			return nil, err
		}

		results = append(results, UploadResult{
			FileName: newName,
			FilePath: savePath,
		})
	}

	return results, nil
}

func (u *Uploader) ValidationFile(file *multipart.FileHeader) error {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowed := false

	for _, a := range u.cfg.AllowedExt {
		if ext == a {
			allowed = true
			break
		}
	}

	if !allowed {
		return errors.New("file extension not allowed")
	}

	maxBytes := u.cfg.MAXSizeMB << 20
	if file.Size > maxBytes {
		return errors.New("file size exceeds limit")
	}

	return nil
}
