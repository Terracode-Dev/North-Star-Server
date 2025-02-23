package hr

import (
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (S *HRService) uploadFile(c echo.Context) error {
	service := c.FormValue("service")
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(500, "file upload issue")
	}

	obj, err := file.Open()
	if err != nil {
		return c.JSON(500, "file Open issue")
	}
	defer obj.Close()

	ext := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + ext

	switch service {
	case "visa":
		err = S.s3.UploadToS3(c.Request().Context(), "nsappvisa", fileName, obj)
		if err != nil {
			return c.JSON(500, "file upload failed")
		}
	case "passport":
		err = S.s3.UploadToS3(c.Request().Context(), "nsappvisa", fileName, obj)
		if err != nil {
			return c.JSON(500, "file upload failed")
		}
	default:
		return c.JSON(500, "invalid service name")
	}

	return c.JSON(201, fileName)
}

type ReqGetFileDownloadUrl struct {
	Service  string `json:"service"`
	FileName string `json:"fileName"`
}

func (S *HRService) getFileDownloadUrl(c echo.Context) error {
	var reqM ReqGetFileDownloadUrl
	if err := c.Bind(&reqM); err != nil {
		return c.JSON(500, "invalid input")
	}

	switch reqM.Service {
	case "visa":
		url, err := S.s3.PresignGetS3Url(c.Request().Context(), "nsappvisa", reqM.FileName)
		if err != nil {
			return c.JSON(500, "file upload failed")
		}
		return c.JSON(200, url)
	case "passport":
		url, err := S.s3.PresignGetS3Url(c.Request().Context(), "nsappvisa", reqM.FileName)
		if err != nil {
			return c.JSON(500, "file upload failed")
		}
		return c.JSON(200, url)
	default:
		return c.JSON(500, "invalid service name")
	}
}
