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
	maxWorkers := 1000
	maxWorkerErrs := 10

	workers := make(chan struct{}, maxWorkers)
	workersErrs := make(chan string, maxWorkerErrs)

	go func() {
		for {
			select {
			case filename := <-workersErrs:
				workers <- struct{}{}
				wg.Add(1)
				log.Println("Retrying file: ", filename)
				go uploadFile(filename, &wg, workers, workersErrs)
			}
		}
	}()

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
		workers <- struct{}{}
		go uploadFile(files[0].Name(), &wg, workers, workersErrs)
	}
	wg.Wait()
}

func uploadFile(filename string, wg *sync.WaitGroup, worker <-chan struct{}, workersErrs chan<- string) error {
	defer wg.Done()

	filepath := "./tmp/" + filename

	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Error opening file %s: %v\n", filepath, err)
		<-worker
		workersErrs <- filename

		return err
	}
	defer file.Close()

	time.Sleep(1 * time.Second)

	log.Printf("Uploaded file %s\n", filename)
	<-worker

	return nil
}
