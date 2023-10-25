package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//postForm

	postParam := url.Values{
		"mobile":      {"xxxxxxx"},
		"isRemberPwd": {"1"},
	}
	resp, err := http.PostForm("http://www.xxx.com/loginRegister/login.do",
		postParam)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
