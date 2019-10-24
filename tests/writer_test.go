package gozip

import (
	"testing"

	"github.com/johnolafenwa/gozip/writer"
)

func TestWriter(t *testing.T) {

	writer, err := writer.New("output.zip")

	if err != nil {
		t.Errorf("%v", err)
		return
	}
	err = writer.AddFile("Cat03.jpg", "Cat03.jpg")
	if err != nil {
		t.Errorf("%v", err)
	}
	err = writer.AddFile("Figure_2.png", "Figure_2.png")
	if err != nil {
		t.Errorf("%v", err)
	}

	err = writer.AddFolder("/home/johnolafenwa/Documents/Kubernetes", "")

	if err != nil {
		t.Errorf("%v", err)
	}

	writer.Save()

}
