# GoZip
An easy to use Go Library for creating and extracting compressed files

![Build Status](https://github.com/johnolafenwa/gozip/workflows/Go/badge.svg "Build Status")


# Installation:
<pre>
go get github.com/johnolafenwa/gozip
</pre>

# Examples

## Writing Zip Files
<pre>
package main

import "github.com/johnolafenwa/gozip/writer"

func main(){

    zipWriter,_ := writer.New("Output.zip")
  
    zipWriter.AddFolder("path-to-folder","")
    
    zipWriter.Save()

}
</pre>


## Extracting Zip Files
<pre>
package main
import "github.com/johnolafenwa/gozip/reader"

func main(){

    zipReader, _ := reader.New("path-to-zip-file")
    
    zipReader.ExtractTo("Destination-Path")
   
    zipReader.Close()
  
}

</pre>


# Functions

## adding a file
The Writer supports adding both individual files and folders to a single archive, you can also set the name property to customize the path in the archive where the file will be stored.

<pre>
package main
import "github.com/johnolafenwa/gozip/writer"

func main(){

    zipWriter,_ := writer.New("MyFiles.zip")
    
    //add a single file
    zipWriter.AddFile("path-to-file/file1.txt","file1.txt")
    
    //add a single file, infer name from path
    zipWriter.AddFile("path-to-file/file2.txt","")
    
    //add a folder
    zipWriter.AddFolder("path-to-folder/FolderOne","FolderOne")
    
    //add a folder, infer name from path
    zipWriter.AddFolder("path-to-folder/FolderTwo","")
    
    zipWriter.Save()

}
</pre>

# Modes
The writer supports two modes for writing zip archives; Store and Deflate. If you want the files compressed, use Store.

<pre>

zipWriter,_ := writer.New("MyFiles.zip")
zipWriter.SetMethod(writer.Deflate)

</pre>

