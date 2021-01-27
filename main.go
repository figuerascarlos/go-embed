package main

import (
	"embed"
	"fmt"
)

func main(){
  //go:embed assets/*
  var f embed.FS
  files, _ := f.ReadDir("assets")
  for _, y := range files {
    fmt.Printf("File %s \n", y.Name())
    content, _ := f.ReadFile(fmt.Sprintf("%s/%s", "assets", y.Name()))
    fmt.Printf("%s", string(content))
  }
  

}
