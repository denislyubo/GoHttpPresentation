package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	bodyRequest := []byte(`name=Denis`)
	bufBody := bytes.NewBuffer(bodyRequest)
	resp, err := http.Post("http://localhost:8080/greeting", "application/x-www-form-urlencoded", bufBody)
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
