package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func calmDown() {
	recover()
}

func main() {
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
	for _, folder := range ways {
		counter(folder)
	}

}
func counter(folder string) {
	files1, err := ioutil.ReadDir(folder)
	if err != nil {
		defer calmDown()
		panic("Cannot read a Dir")
	}
	var fileNames []string
	for _, file := range files1 {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	fmt.Printf("Папка %s содержит файлы: %v\n", folder, fileNames)
}
