package main

import (
	"github.com/JesusKian/WriteUp/src/config"
	"github.com/JesusKian/WriteUp/src/request"
	_ "github.com/JesusKian/WriteUp/src/sql"

	"bufio"
	"os"
)

func main() {
	var (
		scanner *bufio.Scanner  = &bufio.Scanner{}
		file    *os.File        = &os.File{}
		err     error           = nil
	)

	file, err = os.Open("./static/txt/urls.txt")
	if err != nil {
		config.SetLog("E", "main() -> Couldn't Open urls.txt File")
		config.SetLog("D", err.Error())
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		request.GetWriteUps(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		config.SetLog("E", "main() -> Couldn't Read urls.txt Data")
		config.SetLog("D", err.Error())
	}

}
