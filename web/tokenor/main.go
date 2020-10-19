package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var access_token = ""

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
			access_token = r.AccessToken
			fmt.Println("token is: ", access_token)
			time.Sleep(time.Minute * 115)
		}
	}()
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
}

func main() {
	r := gin.Default()
	r.GET("/token", token)
	r.Run(":80")
}

func token(c *gin.Context) {
	c.String(http.StatusOK, access_token)
}
