package main

import (
	"fmt"
	gm "github.com/lgoldstien/goGraphicsMagick"
	"io/ioutil"
	// "os"
	"time"
)

type conversionQueueItem struct {
	src  string
	dest string
}

var conversionQueue = make(chan conversionQueueItem, 50)

var folderPath string

func initConversionWorkers(count int) error {
	var err error = nil
	for w := 1; w <= count; w++ {
		fmt.Println("Starting worker", w, "of", count)
		go conversionWorker(w)
	}
	return err
}

func conversionWorker(id int) error {

	var err error = nil

	for item := range conversionQueue {
		fmt.Printf("Worker", id, "converting job", item.src)
		err = gm.Convert(item.src, item.dest)
	}

	return err
}

func getConversionJobs(path string) error {
	var err error = nil

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {

		conversionQueue <- conversionQueueItem{
			src:  f.Name(),
			dest: folderPath + "/converted/" + f.Name() + ".png",
		}

		fmt.Println(f.Name())
	}

	return err
}

func heartbeat() {
	for {
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	// args := os.Args

	// if len(args) < 1 {
	// 	panic("No folder name given on command line")
	// }

	// folderPath = os.Args[1]

	// err := initConversionWorkers(20)
	// if err != nil {
	// 	panic(err)
	// }

	// err = getConversionJobs(folderPath)
	// if err != nil {
	// 	panic(err)
	// }

	go heartbeat()
}
