package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"web/cloudfunccli"
	"web/oa"

	"github.com/gin-gonic/gin"
)

func init() {
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
			cloudfunccli.AccessToken = r.AccessToken
			fmt.Println("token is: ", cloudfunccli.AccessToken)
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
			cloudfunccli.OATotoken = r.AccessToken
			fmt.Println("official token is: ", cloudfunccli.OATotoken)
			cloudfunccli.DoRequest("consolecommon", `{"tp":"wxe","token":"`+cloudfunccli.OATotoken+`"}`)
			time.Sleep(time.Minute * 115)
		}
	}()
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
}

func main() {
	// oa.BatchSyncOpenids("38_oNKAjSDy6KjfkSFHnS-s2OvHCpxxoIhbUOu3xullXv9HiwsoxKFKek0EpvPjkcULzsochOQNQwR25CifKlxVa_wQt2zHA6f1NWwkPMtI8rIeY2Iz2zjQ3oszWnIVE6oPRtcH_dxjPeTsvt72PWMjABAAHX")
	// return
	r := gin.Default()
	r.GET("/token", token)

	r.GET("/oacallback", oa.Authentication)
	r.POST("/oacallback", oa.OACallback)
	r.Run(":80")
}

func token(c *gin.Context) {
	c.String(http.StatusOK, cloudfunccli.AccessToken)
}
