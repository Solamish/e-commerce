package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type form struct {
	user_id  int
	goods_id int
	shop_id  int
	number   int
}

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		go postTest()
	}
}


func postTest() {
	url := "http://localhost:8080/api/order"

	forms := &form{
		user_id:  64,
		goods_id: 1,
		shop_id:  32,
		number:   10,
	}
	client := &http.Client{Timeout: 5 * time.Second}
	contentType := "application/x-www-form-urlencoded"
	body, err := json.Marshal(forms)
	if err != nil {
		fmt.Println("marshal fail")
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("content-type", contentType)
	if err != nil {
		fmt.Println("post fail")
		return
	}
	defer req.Body.Close()


	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("error :", err1)
	}
	defer resp.Body.Close()

	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("read body fail")
		return
	}
	fmt.Println(string(b))

}


