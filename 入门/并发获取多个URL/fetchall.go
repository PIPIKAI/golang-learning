package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resq, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resq.Body)
	resq.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s %v", url, err)
		return
	}
	sc_time := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f s  %7d  %s", sc_time, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f s elapsed\n", time.Since(start).Seconds())
}
