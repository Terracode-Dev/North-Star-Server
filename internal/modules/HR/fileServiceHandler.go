package hr

import (
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Upload files
// @Description Upload certificate and visa files to S3
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param certificate_file formData file true "Certificate file to upload"
// @Param visa_file formData file true "Visa file to upload"
// @Success 201 {object} Response "File uploaded successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /fileupload [post]
type Response struct {
	Certificate string `json:"certificate"`
	Visa        string `json:"visa"`
}

func (S *HRService) uploadFile(c echo.Context) error {
	certificate, err := c.FormFile("certificate_file")
	if err != nil {
		return c.JSON(500, err.Error())
	}
	visa, err := c.FormFile("visa_file")
	if err != nil {
		return c.JSON(500, err.Error())
	}

	cert_obj, err := certificate.Open()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	defer cert_obj.Close()

	visa_obj, err := visa.Open()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	defer visa_obj.Close()

	cert_ext := filepath.Ext(certificate.Filename)
	visa_ext := filepath.Ext(visa.Filename)
	cert_fileName := uuid.New().String() + cert_ext
	visa_fileName := uuid.New().String() + visa_ext

	err = S.s3.UploadToS3(c.Request().Context(), "nsappcertficates", cert_fileName, cert_obj)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	err = S.s3.UploadToS3(c.Request().Context(), "nsappvisa", visa_fileName, visa_obj)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(201, Response{Certificate: cert_fileName, Visa: visa_fileName})
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
	case "certificate":
		url, err := S.s3.PresignGetS3Url(c.Request().Context(), "nsappvisa", reqM.FileName)
		if err != nil {
			return c.JSON(500, "file upload failed")
		}
		return c.JSON(200, url)
	default:
		return c.JSON(500, "invalid service name")
	}
}
