package writer

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

const Store uint16 = 0
const Deflate uint16 = 8

type Writer struct {
	method      uint16
	filename    string
	zip_writer  *zip.Writer
	output_file *os.File
}

func New(filename string) (*Writer, error) {

	zip_file, err := os.Create(filename)

	if err != nil {
		return nil, err
	}

	return &Writer{zip_writer: zip.NewWriter(zip_file), output_file: zip_file, method: Deflate}, nil

}

func (w *Writer) SetMethod(method uint16) {

	w.method = method

}

func (w *Writer) AddFile(source string, name string) error {

	file, err := os.Open(source)
	defer file.Close()

	if err != nil {
		return err
	}

	file_info, err := file.Stat()

	if err != nil {
		return err
	}

	zip_header, err := zip.FileInfoHeader(file_info)

	if err != nil {
		return err
	}

	zip_header.Name = name
	zip_header.Method = w.method

	writer, err := w.zip_writer.CreateHeader(zip_header)

	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}

func (w *Writer) AddFolder(source string, name string) error {

	err := filepath.Walk(source,
		func(fpath string, info os.FileInfo, err error) error {

			newPath := path.Join(name, fpath[len(source):])

			if err != nil {
				fmt.Println(err)
				return err
			}

			if info.Mode().IsRegular() == true {
				err = w.AddFile(fpath, newPath)

			}

			return err

		})

	return err

}

func (w *Writer) Save() {

	w.zip_writer.Close()

	w.output_file.Close()
}
