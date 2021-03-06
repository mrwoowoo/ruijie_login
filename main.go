package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"regexp"
)

func main() {
	args := os.Args
	if args == nil || len(args) != 3 {
		fmt.Println("使用方法:")
		fmt.Println("程序名", "你的账号", "你的密码")
		fmt.Println("例如:./ruijie_login 123456 123456")
		return
	}
	account := args[1]
	password := args[2]
	resp, err := http.Get("http://www.google.cn/generate_204")
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == 204 {
		log.Println("你已连接网络,无需登录")
	} else {
	resp, err := http.Get("http://www.google.cn/generate_204")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	reg1 := regexp.MustCompile(`'(.*)'`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	sbUrl := reg1.FindAllStringSubmatch(sb, -1)[0][1]
	sbUrlSlice := strings.Split(sbUrl, "?")
	loginURL := strings.Replace(sbUrlSlice[0], "index.jsp", "InterFace.do?method=login", -1)
	queryString := sbUrlSlice[1]
	data := "userId=" + account + "&password=" + password +"&service=&queryString=" + queryString + "&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false"
	responseBody := strings.NewReader(data)
	resp, err = http.Post(loginURL, "application/x-www-form-urlencoded; charset=UTF-8", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb = string(body)
	log.Println(sb)
	}
}