package main

import (
	"fmt"

	"johnolafenwa.io/gozip/writer"
)

func main() {

	writer, err := writer.New("kube2.tar", writer.TAR)

	if err != nil {
		fmt.Println(err)
		return
	}

	writer.AddFolder("/home/johnolafenwa/Documents/Kubernetes", "Kubernetes")

	writer.Save()

}
