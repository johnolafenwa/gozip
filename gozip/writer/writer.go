package writer

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//Store mode archives files but does not compress them
const Store uint16 = 0

//Deflate mode compresses archived files
const Deflate uint16 = 8

type Writer struct {
	method      uint16
	filename    string
	zip_writer  *zip.Writer
	output_file *os.File
}

//New eturns an instance of the Writer, the filename is used to create the resulting zip file
func New(filename string) (*Writer, error) {

	zip_file, err := os.Create(filename)

	if err != nil {
		return nil, err
	}

	return &Writer{zip_writer: zip.NewWriter(zip_file), output_file: zip_file, method: Deflate}, nil

}

//SetMethod sets the method to be used for compression. This can Store or Deflate from Writer constants
func (w *Writer) SetMethod(method uint16) {

	w.method = method

}

//AddFile creates a new file entry in the zip archive. source specifies an existing file to read contents from. The name is the path name of the file
//in the archive. Note that if name is an empty string, the filename from source will be used and the file will be created in the root path.
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

	if strings.TrimSpace(name) == "" {

		zip_header.Name = file_info.Name()

	} else {
		zip_header.Name = name
	}
	zip_header.Method = w.method

	writer, err := w.zip_writer.CreateHeader(zip_header)

	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}

//AddFolder creates a new folder entry int the archive and recursively adds all files within the folder. The file structure of the folder
//is preserved. The name property is used to name the folder path, however, if the name is empty, the folder name will be inferred from
//the source path and the folder will be stored in the root directory.
func (w *Writer) AddFolder(source string, name string) error {

	if strings.TrimSpace(name) == "" {
		name = filepath.Base(source)
	}

	err := filepath.Walk(source,
		func(fpath string, info os.FileInfo, err error) error {

			newPath := path.Join(name, fpath[len(source):])

			if err != nil {
				return err
			}

			if info.Mode().IsRegular() == true {
				err = w.AddFile(fpath, newPath)

			}

			return err

		})

	return err

}

//Save closes the underlying zip file and zip writer
func (w *Writer) Save() {

	w.zip_writer.Close()

	w.output_file.Close()
}
