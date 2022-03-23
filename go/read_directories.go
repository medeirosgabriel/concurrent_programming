package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
	channel := make(chan string)
	finish := make(chan int) 
    go walk("./directory/", channel)
	go readFile(channel, finish)
	for range finish {}
}

func walk(path string, ch chan string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
        log.Fatal(err)
    }
	
	var directories []string

    for _, f := range files {
		if (!f.IsDir()) {
			file_path := path + f.Name()
			ch <- file_path
		} else {
			directories = append(directories, f.Name())	
		}
    }
	if (len(directories) > 0) {
		for _, element := range directories {
			walk(path + element + "/", ch)
		}
	} else{
		close(ch)
	}
}

func readFile(ch chan string, finish chan int) {
	for {
		path, ok := <-ch
		if (ok) {
			body, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			data := string(body)
			firstCharacter := string([]rune(data)[1])
			bytes := []rune(data)[1]
			if (bytes % 2 == 0) {
				fmt.Printf("%s - %s - %d\n", path, firstCharacter, bytes)
			}
		} else {
			break
		}
	}
	close(finish)
}
