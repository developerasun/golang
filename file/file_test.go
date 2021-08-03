package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "C:/Users/nello/Downloads/myPandas.py"
	file, _ := os.Open(path)

	defer file.Close()

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	// filepath.Base: Base returns the last element of path.
	// Trailing path separators are removed before extracting the last element.
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))

	assert.Nil(err)

	io.Copy(multi, file)

	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)

	// Header contains the request header fields either received by the server or to be sent by the client.
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadsHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err2 := os.Stat(uploadFilePath)

	assert.NoError(err2)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)

	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)

}
