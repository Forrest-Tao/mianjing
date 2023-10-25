package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//post
	resp, err := http.Post("http://www.xxx.com/loginRegister/login.do",
		"application/x-www-form-urlencoded", strings.NewReader("mobile=xxxxxx&isRemeberPwd=1"))
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
