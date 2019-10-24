package tests

import (
	"testing"

	"github.com/johnolafenwa/gozip/reader"
)

func TestReader(t *testing.T) {

	read, err := reader.New("output.zip")

	if err != nil {
		t.Errorf("The following error occured %v", err)

		return
	}

	err = read.ExtractTo("output")

	if err != nil {
		t.Errorf("The following error occured %v", err)

	}

}
