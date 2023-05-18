package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFileTrainer(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Save file to disk
	if err := c.SaveFile(file, fmt.Sprintf("storage/trainer/%s", filename)); err != nil {
		return "", err
	}

	// Create img url for database
	path := fmt.Sprintf("%s/images/trainer/%s", os.Getenv("BASE_URL"), filename)

	return path, nil
}

func SaveFileEquipment(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Save file to disk
	if err := c.SaveFile(file, fmt.Sprintf("storage/equipment/%s", filename)); err != nil {
		return "", err
	}

	// Create img url for database
	path := fmt.Sprintf("%s/images/equipment/%s", os.Getenv("BASE_URL"), filename)

	return path, nil
}
