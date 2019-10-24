package tests

import (
	"testing"

	"github.com/johnolafenwa/gozip/writer"
)

func TestWriter(t *testing.T) {

	writer, err := writer.New("output.zip")

	if err != nil {
		t.Errorf("The following error occured %f", err)
	}
	err = writer.AddFile("Cat03.jpg", "Cat03.jpg")
	if err != nil {
		t.Errorf("The following error occured %f", err)
	}
	err = writer.AddFile("Figure_2.png", "Figure_2.png")
	if err != nil {
		t.Errorf("The following error occured %f", err)
	}

	writer.AddFolder("/home/johnolafenwa/Documents/Kubernetes", "")

	writer.Save()

}
