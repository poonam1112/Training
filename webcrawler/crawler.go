package main

import ( 
"encoding/json"       
  "fmt"
  "net/http"
  "io"      
  "io/ioutil"     
  "time"
)


//creating lang type struct for webcrawler
type Lang struct {
  Name string
  URL string
  Bytes int64
  Time time.Duration

}

//Calculating Bytes and time for webcrawlers
func retrieve(uri string) ( int64, time.Duration) {  
                           
  now := time.Now()// buil-in function used for time
  resp1, err := http.Get(uri)// fetching url
  if err != nil { 
    t := time.Since(now)           
    return 0, t                
  }

  num,_ :=io.Copy(ioutil.Discard, resp1.Body)//built-in function for byte count
  t := time.Since(now)
  return num, t
}

func pfunc(l Lang) {
  //applying JSON for Printing Struct elements in a format
      res1D :=&Lang{
      	Name:  l.Name,
        URL: l.URL ,
        Bytes: l.Bytes,
        Time: l.Time,
}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
    //fmt.Println("BYTES ",l.bytes)
}

//crawl function used to retrieve url, Bytes, time
func crawl(pfunc func(Lang), lang Lang){
  name:=[3]string{"python","ruby","golang"}
  url:= [3]string{"https://www.python.org/", "https://www.ruby-lang.org/en/","https://golang.org/"}

var Total_time time.Duration
  for i,url_val := range url{
    lang.URL = url_val
    lang.Name=name[i]
    lang.Bytes,lang.Time = retrieve(url_val)
    Total_time = Total_time+lang.Time	//total time for fetching all sites
    pfunc(lang) //calling pfunc for printing

    //fmt.Println("TIME: ",lang.time)
}

  fmt.Println("\nTOTAL TIME: ",Total_time)
}

func main() {
	//main function calling crawl
  var l Lang
  crawl(pfunc,l)                   
}