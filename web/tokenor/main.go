package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var access_token = ""
var wxeToken = ""

func init1() {
	go func() {
		for {
			resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx9b588b2b3f090400&secret=08d530ef7af12573a04586491aff47a7")
			if err != nil {
				fmt.Println(err)
				time.Sleep(time.Minute)
				continue
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				time.Sleep(time.Minute)
				continue
			}

			r := &TokenResp{}

			if err := json.Unmarshal(data, r); err != nil {
				fmt.Println(err)
				time.Sleep(time.Minute)
				continue
			}
			fmt.Println(string(data))
			access_token = r.AccessToken
			fmt.Println("token is: ", access_token)
			time.Sleep(time.Minute * 115)
		}
	}()

	time.Sleep(time.Second * 5)
	// 公众号
	go func() {
		for {
			resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxe66de2494b1702e5&secret=9c1ff5385ad8c8e50c15dd3be9842137")
			if err != nil {
				fmt.Println(err)
				time.Sleep(time.Minute)
				continue
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				time.Sleep(time.Minute)
				continue
			}

			r := &TokenResp{}

			if err := json.Unmarshal(data, r); err != nil {
				fmt.Println(err)
				time.Sleep(time.Minute)
				continue
			}
			fmt.Println(string(data))
			wxeToken = r.AccessToken
			fmt.Println("wxe token is: ", access_token)
			doRequest("consolecommon", `{"tp":"wxe","token":"`+wxeToken+`"}`)
			time.Sleep(time.Minute * 115)
		}
	}()
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
}

func main() {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/user/get?access_token=38_g477FAmfZkj00VwMngC8Bj-_7d5dV4WPxy7VpMbKHO5WsT5sYMrJFM2nl0bbj_OY1tBxWEbL5xacp2arSUyrizhTbf5tIj5JzWIq3v8pO-f9barQXoqedDn5XrkmMXXbsKIjd7fJ_RD67iU1GDEcADAURS")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(data))

	rr := R{}
	if err := json.Unmarshal(data, &rr); err != nil {
		panic(err)
	}
	// fmt.Println(rr.Data)

	for i, v := range rr.Data.OpenID {
		fmt.Println("executing...", i)
		resp1, err := http.Get("https://api.weixin.qq.com/cgi-bin/user/info?access_token=38_g477FAmfZkj00VwMngC8Bj-_7d5dV4WPxy7VpMbKHO5WsT5sYMrJFM2nl0bbj_OY1tBxWEbL5xacp2arSUyrizhTbf5tIj5JzWIq3v8pO-f9barQXoqedDn5XrkmMXXbsKIjd7fJ_RD67iU1GDEcADAURS&openid=" + v + "&lang=zh_CN")
		if err != nil {
			panic(err)
		}
		data, err = ioutil.ReadAll(resp1.Body)
		if err != nil {
			panic(err)
		}
		// fmt.Println(string(data))
		r1 := &R1{}
		if err := json.Unmarshal(data, &r1); err != nil {
			panic(err)
		}
		if r1.NickName == "易" {
			fmt.Println(string(data))
		}
	}

	return
	r := gin.Default()
	r.GET("/token", token)
	r.Run(":80")
}

type R1 struct {
	OpenID   string `json:"openid"`
	NickName string `json:"nickname"`
}

type R struct {
	Total int64 `json:"total"`
	Count int64 `json:"count"`
	Data  *D    `json:"data"`
}

type D struct {
	OpenID []string `json:"openid"`
}

func token(c *gin.Context) {
	c.String(http.StatusOK, access_token)
}

const (
	// env = "dev-osmu3"
	env      = "release-8tcge"
	baseURL  = "https://api.weixin.qq.com/tcb/invokecloudfunction"
	funcName = "console"
)

func doRequest(funcName string, params string) (string, error) {
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, funcName), "application/json", strings.NewReader(params))
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))
	return string(data), nil
}
