package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/greeting")
	if err != nil {
		fmt.Printf("error: %+v", err)
		return
	}
	fmt.Println(resp.Status)
	for k, v := range resp.Header {
		fmt.Printf("%s: %+v\n", k, v)
	}
	defer resp.Body.Close()
	buf := bufio.NewScanner(resp.Body)
	for buf.Scan() {
		fmt.Println(buf.Text())
	}
}
