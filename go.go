package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

var maxChannelBuffer = 6
var ch = make(chan int32, maxChannelBuffer)

func main() {
	fileName := flag.String("f", "", "File name which contains download list")
	concurrentCount := flag.Int("c", 4, "*Concurrent files to be downloaded")
	currentWorkingDirectory, err := os.Getwd()
	baseDestination := flag.String("d", currentWorkingDirectory, "Destination folder for files/folder")
	flag.Parse()

	if len(*fileName) == 0 {
		log.Fatal("File name required")
	}

	maxChannelBuffer = *concurrentCount
	ch = make(chan int32, maxChannelBuffer)
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		urlString := scanner.Text()
		v, err := url.Parse(urlString)
		if err != nil {
			continue
		}
		count++

		dirName := path.Join(*baseDestination, filepath.Dir(v.Path))
		err = os.MkdirAll(dirName, os.ModeDir)
		if err != nil {
			log.Println(err)
		}
		go DownloadFile(path.Join(dirName, filepath.Base(v.Path)), urlString)
		if count == maxChannelBuffer-1 {
			<-ch
			count--
		}
	}

	//Just in case some downloads pending
	for count > 0 {
		<-ch
		count--
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)

		ch <- 1
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	log.Println(filepath)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Println(err)

		ch <- 1
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println(err)
	}
	ch <- 2
	return err
}
