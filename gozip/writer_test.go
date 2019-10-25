package tests

import (
	"testing"

	"github.com/johnolafenwa/gozip/writer"
)

func TestWriter(t *testing.T) {

	writer, err := writer.New("test.zip")

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if err != nil {
		t.Errorf("%v", err)
	}
	err = writer.AddFile("../testfiles/walk.jpg", "walk.jpg")
	if err != nil {
		t.Errorf("%v", err)
	}
	err = writer.AddFile("../testfiles/jaguar.jpg", "jaguar.jpg")
	if err != nil {
		t.Errorf("%v", err)
	}

	err = writer.AddFolder("../testfiles/Pictures", "Pictures")

	if err != nil {
		t.Errorf("%v", err)
	}

	writer.Save()

}
