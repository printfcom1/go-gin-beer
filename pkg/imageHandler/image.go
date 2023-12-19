package imageHandler

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func CreateImage(name string, c *gin.Context) (*string, error) {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	err = c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return nil, err
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileName := uuid.New().String() + "_" + name + ".png"
	filePath := envMap["FILE_PATH_IMAGE"]
	savePath := filepath.Join(filePath, fileName)
	err = c.SaveUploadedFile(header, savePath)
	if err != nil {
		return nil, err
	}

	return &fileName, nil
}

func DeleteImage(file string) error {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		return err
	}
	Path := envMap["FILE_PATH_IMAGE"]
	filePath := filepath.Join(Path, file)
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
