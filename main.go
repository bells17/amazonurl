package main

import (
	"fmt"
	"os"
	"regexp"
    "errors"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("URL is not found")
		os.Exit(1)
	}
    url, err := url(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    fmt.Println(url)
}

func url(url string) (string, error) {
    reg := regexp.MustCompile(`\/(dp\/|gp\/product\/).{10}`)
    ret := reg.FindStringSubmatch(url)
    if ret == nil {
        return "", errors.New("URL is not matched")
    }
    return fmt.Sprint("http://www.amazon.co.jp", ret[0]), nil
}
