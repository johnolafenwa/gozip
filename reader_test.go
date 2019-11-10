package gozip

import (
	"testing"

	"github.com/johnolafenwa/gozip/reader"
)

func TestReader(t *testing.T) {

	read, err := reader.New("testfiles/Images.zip")

	if err != nil {
		t.Errorf("%v", err)

		return
	}

	err = read.ExtractTo("output")

	if err != nil {
		t.Errorf("%v", err)

	}

}
