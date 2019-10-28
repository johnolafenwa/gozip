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

//New returns an instanced of the Reader and opens the zipfile specified as source for reading.
func New(source string) (*Reader, error) {

	zip_reader, err := zip.OpenReader(source)

	if err != nil {
		return nil, err
	}

	return &Reader{zip_reader: zip_reader}, nil
}

//ExtractTo extracts all of the files and folders in the zip archive to the the directory specified by dest. The filel strcuture is preserved.
func (r *Reader) ExtractTo(dest string) error {

	for _, file := range r.zip_reader.File {

		filePath := path.Join(dest, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, file.FileInfo().Mode().Perm())

			continue
		}

		data, err := file.Open()

		if err != nil {
			return err
		}

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

//Files returns a list of files in the zip archive
func (r *Reader) Files() []*zip.File {

	return r.zip_reader.File
}

//Close closes the underlying zip reader
func (r *Reader) Close() {
	r.zip_reader.Close()
}
