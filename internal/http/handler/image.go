package handler

import (
	"bytes"
	"io"
	"net/http"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/constant"
	"github.com/derangga/shopifyx/internal/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	maxFileSize = 2 << 20  // 2MB
	minFileSize = 10 << 10 // 10KB
)

// ImageHandler is a struct to handle http request in image handler
type ImageHandler struct {
	imageUC  internal.ImageUsecase
	validate *validator.Validate
}

// NewImageHandler is a function to initialize image handler
func NewImageHandler(imageUC internal.ImageUsecase, validate *validator.Validate) *ImageHandler {
	return &ImageHandler{
		imageUC:  imageUC,
		validate: validate,
	}
}

func (h *ImageHandler) Upload(c echo.Context) error {
	src, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	fileType := src.Header.Get(echo.HeaderContentType)
	if fileType != constant.MIMETypeJPEG {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	if src.Size > maxFileSize || src.Size < minFileSize {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	file, err := src.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}
	defer file.Close()

	// Create a buffer to store the file contents
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	res, err := h.imageUC.Upload(c.Request().Context(), buffer)
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response.Image{
		URL: res,
	})
}
