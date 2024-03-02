package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func calmDown() {
	recover()
}

func main() {
	start := time.Now()
	p := filepath.Join(os.Args[1])
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}

	var ways []string
	for _, file := range files {
		a := file.Name()
		path := filepath.Join(p, a)
		ways = append(ways, path)
	}

	var wg sync.WaitGroup
	for _, folder := range ways {
		wg.Add(1)
		go counter(folder, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Время выполнения: %s\n", elapsed)
}

func counter(folder string, wg *sync.WaitGroup) {
	defer wg.Done()

	files1, err := ioutil.ReadDir(folder)
	if err != nil {
		defer calmDown()
		panic("oh no")
	}

	var fileNames []string
	for _, file := range files1 {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	fmt.Printf("Папка %s содержит файлы: %v\n", folder, fileNames)
}
