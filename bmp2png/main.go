package main

import (
    "golang.org/x/image/bmp"
    "os"
    "log"
    "image/png"
)

func main ()  {
   // file,err:=os.OpenFile("girl.bmp")
    file, err := os.Open("girl.bmp") // For read access.
    if err != nil {
        log.Fatal(err)
    }
    im,err:=bmp.Decode(file)
    if err!=nil {
        log.Fatal(err)
    }
    file1,err:=os.Create("girl.png")
    if err != nil {
        log.Fatal(err)
    }
    err=png.Encode(file1,im)
    if err != nil {
        log.Fatal(err)
    }
    
}