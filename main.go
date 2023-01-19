package main

import (
	"github.com/JesusKian/WriteUp/src/config"
	// "github.com/JesusKian/WriteUp/src/sql"
	"github.com/JesusKian/WriteUp/src/request"
	// "github.com/JesusKian/WriteUp/src/structure"

	"bufio"
	"os"
	"sync"
	// "encoding/xml"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
)

func init() {
	config.ConnectToSqlDatabase()
	config.DatabaseStatus()
}

func main() {
	var (
		err     error           = nil
		file    *os.File        = &os.File{}
		wg      *sync.WaitGroup = &sync.WaitGroup{}
		scanner *bufio.Scanner  = &bufio.Scanner{}
	)
	file, err = os.Open("./static/txt/urls.txt")
	if err != nil {

	}
	defer file.Close()
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		go func(_url string) {
			request.GetWriteUps(_url)
			wg.Done()
		}(scanner.Text())
	}
	wg.Wait()
	// rss := structure.Rss{}
	// res, err := http.Get("https://medium.com/feed/tag/ctf")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// xml.Unmarshal(data, &rss)
	// for i, j := range rss.Channel[0].Items {
	// 	fmt.Println(i)
	// 	fmt.Println(j.Title)
	// 	fmt.Println()
	// }
}
