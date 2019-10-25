# GoZip
An easy to use Go Library for creating and extracting compressed files

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

    writer,_ := writer.New("Output.zip")
  
    writer.AddFolder("path-to-folder")
    
    writer.Save()

}
</pre>


## Extracting Zip Files
<pre>
package main
import "github.com/johnolafenwa/gozip/reader"

func main(){

    reader, _ := reader.New("path-to-zip-file")
    
    reader.ExtractTo("Destination-Path")
    
    reader.Close()
  
}

</pre>
