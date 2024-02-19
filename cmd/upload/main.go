package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
  folder, err := os.Open("./tmp")
  if err != nil {
    log.Panic(err)
  }
  defer folder.Close()

  wg := sync.WaitGroup{}

  for {
    files, err := folder.ReadDir(1)
    if err != nil {
      if err == io.EOF {
        break
      }

      log.Printf("Error reading folder: %s\n", err)
      continue
    }


    wg.Add(1)
    go uploadFile(files[0].Name(), &wg)
  }
  wg.Wait()
}

func uploadFile(filename string, wg *sync.WaitGroup) error {
  defer wg.Done()

  filepath := "./tmp/" + filename
  
  file, err := os.Open(filepath)
  if err != nil {
    log.Printf("Error opening file %s: %v\n", filepath, err)
  }
  defer file.Close()

  time.Sleep(1 * time.Second)

  log.Printf("Uploaded file %s\n", filename)

  return nil
}
