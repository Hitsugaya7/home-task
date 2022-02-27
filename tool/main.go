package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	var parallel int
	flag.IntVar(&parallel, "parallel", 10, "Number of parallel requests to be executed")
	flag.Parse()
	args := flag.Args()
	if parallel > len(args) {
		parallel = len(args)
	}
	results := ParallelGetRequests(args, parallel)
	for _, item := range results {
		fmt.Println(item)
	}
}

func ParallelGetRequests(urls []string, concurrencyLimit int) []string {
	semaphore := make(chan struct{}, concurrencyLimit)
	resultsChan := make(chan string)
	defer func() {
		close(semaphore)
		close(resultsChan)
	}()
	for _, url := range urls {
		go func(url string) {
			if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
				url = "http://" + url
			}
			semaphore <- struct{}{}
			client := http.Client{}
			resp, err := client.Get(url)
			if err != nil {
				log.Fatal("Request Error: ", err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("Error reading the body: ", err)
			}
			resultsChan <- fmt.Sprintf("%s %s", url, MD5Hash(body))
			<-semaphore

		}(url)
	}
	var results []string
	for {
		result := <-resultsChan
		results = append(results, result)
		if len(results) == len(urls) {
			break
		}
	}
	return results
}

func MD5Hash(message []byte) string {
	hash := md5.Sum(message)
	return hex.EncodeToString(hash[:])
}
