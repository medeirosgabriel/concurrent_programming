package main

import (
    "fmt"
    "io/ioutil"
    "sync"
    "log"
)

func main() {
	channel := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go walk("./directory/", channel, 0)
	go readFile(channel, &wg)
	wg.Wait()
}

func walk(path string, ch chan string, aux int) {
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
			walk(path + element + "/", ch, aux + 1)
		}
	}

	if (aux == 0) {
		close(ch)
	}
}

func readFile(ch chan string, wg *sync.WaitGroup) {
	for path := range ch {
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
	}
	wg.Done()
}
