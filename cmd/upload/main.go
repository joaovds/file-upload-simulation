package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
  folder, err := os.Open("./tmp")
  if err != nil {
    log.Panic(err)
  }
  defer folder.Close()

  for {
    files, err := folder.ReadDir(1)
    if err != nil {
      if err == io.EOF {
        break
      }

      log.Printf("Error reading folder: %s\n", err)
      continue
    }

    uploadFile(files[0].Name())
  }
}

func uploadFile(filename string) error {
  log.Println("Uploading file: ", filename)

  time.Sleep(1 * time.Second)

  return nil
}
