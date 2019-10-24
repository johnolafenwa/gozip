package reader

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
)

type Reader struct {
	zip_reader *zip.ReadCloser
}

func New(source string) (*Reader, error) {

	zip_reader, err := zip.OpenReader(source)

	if err != nil {
		return nil, err
	}

	return &Reader{zip_reader: zip_reader}, nil
}

func (r *Reader) ExtractTo(dest string) error {

	for _, file := range r.zip_reader.File {

		data, err := file.Open()

		if err != nil {
			return err
		}

		filePath := path.Join(dest, file.Name)

		err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)

		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())

		if err != nil {
			return err
		}

		io.Copy(outFile, data)

		outFile.Close()
	}

	return nil

}

func (r *Reader) Files() []*zip.File {

	return r.zip_reader.File
}
