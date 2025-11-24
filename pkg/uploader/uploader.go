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
	Destination     string
	MaxSingleSize   int64
	MaxMultipleSize int64
	AllowedExt      []string
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

func (u *Uploader) UploadSingle(c *fiber.Ctx, field string) (*UploadResult, error) {
	file, err := c.FormFile(field)
	if err != nil {
		return nil, err
	}

	if file.Size > u.cfg.MaxSingleSize {
		return nil, fmt.Errorf("file size is too large, maximum %d mb", u.cfg.MaxSingleSize/1024/1024)
	}

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

func (u *Uploader) UploadMultiple(c *fiber.Ctx, field string) ([]UploadResult, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File[field]
	if files == nil {
		return nil, err
	}

	var results []UploadResult
	var totalSize int64 = 0

	for _, file := range files {
		totalSize += file.Size

		if totalSize > u.cfg.MaxMultipleSize {
			return nil, fmt.Errorf("total file size is too large, maximum %d mb", u.cfg.MaxMultipleSize/1024/1024)
		}

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

	return nil
}
