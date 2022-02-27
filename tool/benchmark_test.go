package main

import (
	"fmt"
	"testing"
)

var urls []string

func TestMain(m *testing.M) {
	for i := 0; i < 10; i++ {
		urls = append(urls, "adjust.com")
		urls = append(urls, "google.com")
		urls = append(urls, "facebook.com")
		urls = append(urls, "yahoo.com")
		urls = append(urls, "yandex.com")
		urls = append(urls, "twitter.com")
		urls = append(urls, "reddit.com/r/funny")
		urls = append(urls, "reddit.com/r/notfunny")
		urls = append(urls, "baroquemusiclibrary.com")
	}
	m.Run()

}

func BenchmarkTestParallelGetRequests_20(b *testing.B) {
	requests := ParallelGetRequests(urls, 3)
	fmt.Println(len(requests))
}

func BenchmarkTestParallelGetRequests_30(b *testing.B) {

	requests := ParallelGetRequests(urls, 4)
	fmt.Println(len(requests))
}

func BenchmarkTestParallelGetRequests_50(b *testing.B) {
	requests := ParallelGetRequests(urls, 5)
	fmt.Println(len(requests))
}
