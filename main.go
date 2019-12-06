package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
)

func main() {
	file := flag.String("f", "", "input file, uses stdin if left empty")
	flag.Parse()
	var f io.ReadCloser
	if *file == "" {
		f = os.Stdin
	} else {
		urls, err := os.Open(*file)
		if err != nil {
			log.Fatal(err)
		}
		f = urls
	}
	defer f.Close()
	params := make(map[string]struct{})
	s := bufio.NewScanner(f)
	for s.Scan() {
		u, err := url.Parse(s.Text())
		if err != nil {
			continue
		}
		for k := range u.Query() {
			params[k] = struct{}{}
		}
	}
	for k := range params {
		fmt.Println(k)
	}
}
