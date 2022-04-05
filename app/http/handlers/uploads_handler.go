package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type UploadsHandler struct {
}

func NewUploadsHandler() *UploadsHandler {
	return &UploadsHandler{}
}

func (u *UploadsHandler) GetFileById(context *gin.Context) {
	id := context.Param("id")

	if len(id) == 0 {
		context.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	path := "storage/uploads/" + id

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		context.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		context.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	context.Writer.WriteHeader(http.StatusOK)
	context.Writer.Header().Set("Content-Type", "application/octet-stream")
	context.Writer.Write(fileBytes)
}
