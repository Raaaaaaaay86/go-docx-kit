package docx

import (
	"archive/zip"
	"io"
)

type DocxZipFile struct {
	Name string
	Data []byte
}

func NewDocxZipFile(zip *zip.File) (*DocxZipFile, error) {
	reader, err := zip.Open()
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return &DocxZipFile{Name: zip.Name, Data: data}, nil
}