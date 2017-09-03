package utils

import (
	"./surfer"
	"encoding/json"
	logs "github.com/yangaowei/gologs"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GetHtml(req surfer.Request) (resp string, err error) {
	log.Println("get html from url ", req.GetUrl())
	down := surfer.New()
	if response, e := down.Download(req); e == nil {
		defer response.Body.Close()
		bytes, _ := ioutil.ReadAll(response.Body)
		resp = string(bytes)
	} else {
		//log.Println("err:", e)
		resp = "resp"
		err = e
	}
	return resp, err
}

func GetContent(url string, data map[string]interface{}) (resp string, err error) {
	request := &surfer.DefaultRequest{Url: url, TryTimes: 3}
	if value, ok := data["proxy"]; ok {
		request.Proxy = value.(string)
	}
	if header, ok := data["header"]; ok {
		logs.Log.Debug("header %v", data["header"])
		request.Header = header.(http.Header)
	}
	request.GetUrl()
	return GetHtml(request)
}

//正则表达式相关内容
func MatchString(pattern string, content string) (b bool) {
	if m, _ := regexp.MatchString(pattern, content); !m {
		return false
	}
	return true
}

func RxOf(pattern string, content string, index int) (rcontent string) {
	re, _ := regexp.Compile(pattern)
	submatch := re.FindStringSubmatch(content)
	for i, v := range submatch {
		if i == index {
			rcontent = v
			break
		}
	}
	return
}

func R1(pattern string, content string) (rcontent string) {
	return RxOf(pattern, content, 1)
}

func R1Of(patterns []string, s string) (rcontent string) {
	for _, pattern := range patterns {
		if rcontent = R1(pattern, s); len(rcontent) > 0 {
			break
		}
	}
	return
}

func FindSubAll(pattern string, content string) (rcontent []string) {
	re, _ := regexp.Compile(pattern)
	rcontent = re.FindStringSubmatch(content)
	return
}

func FindAll(pattern string, content string) (rcontent []string) {
	re, _ := regexp.Compile(pattern)
	rcontent = re.FindAllString(content, 100)
	return
}

//json

func Loads(sjson string) (result interface{}) {
	err := json.Unmarshal([]byte(sjson), &result)
	if err != nil {
		log.Println("error:", err)
	}
	return
}
