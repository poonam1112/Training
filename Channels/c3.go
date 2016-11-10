package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//creating lang type struct for webcrawler
type Lang struct {
	Name  string
	URL   string
	Bytes int64
	Time  time.Duration
}

//Calculating Bytes and time for webcrawlers
func retrieve(uri string) (int64, time.Duration) {

	now := time.Now()           // buil-in function used for time
	resp1, err := http.Get(uri) // fetching url
	if err != nil {
		t := time.Since(now)
		return 0, t
	}

	num, _ := io.Copy(ioutil.Discard, resp1.Body) //built-in function for byte count
	t := time.Since(now)
	return num, t

}

func pfunc(ch chan Lang) {
	//Printing in Golang Format

	//applying JSON for Printing Struct elements in a format

	l := <-ch
	fmt.Println("GoLang Format: %v", l)

	fmt.Println("JSON Format")
	res1D := &Lang{
		Name:  l.Name,
		URL:   l.URL,
		Bytes: l.Bytes,
		Time:  l.Time,
	}
	res1B, _ := json.Marshal(res1D) //Printing JSON format using json.Marshal
	fmt.Println(string(res1B))
	fmt.Println("\n")

}

//crawl function used to retrieve url, Bytes, time
func crawl(pfunc func(chan Lang), lang Lang, w *sync.WaitGroup) {
	name := [3]string{"python", "ruby", "golang"}
	url := [3]string{"https://www.python.org/", "https://www.ruby-lang.org/en/", "https://golang.org/"}
	ch := make(chan Lang, 3)
	var Cum_time time.Duration
	var tbytes int64
	Tnow := time.Now()
	for i, url_val := range url {
		w.Add(1)
		go routine(name[i], url_val, w, ch, &Cum_time, &tbytes) //Go routine
	}

	w.Wait()
	total := time.Since(Tnow)
	fmt.Println("\nTOTAL TIME con.:", total)
	fmt.Println("TOTAL cumulative TIME : %v", Cum_time)
	fmt.Println("TOTAL bytes read form all 3 sites %v", tbytes)
}

func routine(name string, url string, w *sync.WaitGroup, ch chan Lang, Cum_time *time.Duration, tbytes *int64) {

	var lang Lang
	lang.URL = url
	lang.Name = name
	lang.Bytes, lang.Time = retrieve(url)
	ch <- lang
	pfunc(ch)
	*Cum_time = *Cum_time + lang.Time
	*tbytes = *tbytes + lang.Bytes
	w.Done()

}

func main() {
	wg := sync.WaitGroup{}
	//main function calling crawl
	var l Lang

	crawl(pfunc, l, &wg)

	wg.Wait()

}
