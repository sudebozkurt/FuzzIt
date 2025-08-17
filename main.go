package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"log"
)

func main() {

	asciiArt()

	speedPtr := flag.Int("s", 72, "Tarama hızı")
	txtPtr := flag.String("txt", "wordlist3.txt", "Wordlist dosyası")
	urlPtr := flag.String("u", "http://localhost", "Hedef URL")
	flag.Parse()

	wordlist, err := ioutil.ReadFile(*txtPtr)
	if err != nil {
		fmt.Println("Wordlist dosyası okunamadı:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(wordlist), "\n")
	var wg sync.WaitGroup

	fuzzer := func(line string) {
		defer wg.Done()

		line = strings.TrimRight(line, "\r")

		client := http.Client{
			Timeout: time.Duration(5 * time.Second),
		}

		targetURL := fmt.Sprintf("%s/%s", *urlPtr, line)

		resp, err := client.Get(targetURL)
		if err != nil {
			fmt.Printf("Hata: %s\n", err)
			return
		}

		defer resp.Body.Close()
		fmt.Printf("URL: %s, Status: %s\n", targetURL, resp.Status)
	}

	rateLimiter := time.Tick(time.Second / time.Duration(*speedPtr))

	for _, line := range lines {
		wg.Add(1)
		go fuzzer(line)
		<-rateLimiter
	}

	wg.Wait()
}

func asciiArt(){
	filePath := "asciiart.txt"

	asciiArt, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("ASCII sanat dosyasını okuma hatası:", err)
	}

	fmt.Println(string(asciiArt))
	fmt.Println()
	fmt.Println()
}
