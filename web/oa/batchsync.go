package oa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"web/cloudfunccli"

	"github.com/opentracing/opentracing-go/log"
)

func BatchSyncOpenids(accessToken string) error {

	rr, err := ListOAUsers(accessToken)
	if err != nil {
		return err
	}

	for i, openid := range rr.Data.OpenID {
		fmt.Println("executing...", i)
		info, err := GetOAUserInfo(accessToken, openid)
		if err != nil {
			log.Error(err)
			continue
		}
		info.Tp = "oacallback"
		jsonData, err := json.Marshal(info)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(jsonData))
		cloudfunccli.DoRequest("consolecommon", string(jsonData))
	}

	return nil
}

// official accounts user list resp
type R struct {
	Total int64 `json:"total"`
	Count int64 `json:"count"`
	Data  *D    `json:"data"`
}

// official accounts user list openids data
type D struct {
	OpenID []string `json:"openid"`
}

func ListOAUsers(accessToken string) (*R, error) {
	listUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s", accessToken)
	resp, err := http.Get(listUrl)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rr := &R{}
	if err := json.Unmarshal(data, rr); err != nil {
		return nil, err
	}
	return rr, nil
}

// official accounts user info
type OAUserInfo struct {
	OpenID   string `json:"openid"`
	NickName string `json:"nickname"`
	UnionID  string `json:"unionid"`
	Tp       string `json:"tp"`
}

func GetOAUserInfo(accessToken string, openid string) (*OAUserInfo, error) {
	userInfoUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN", accessToken, openid)
	resp, err := http.Get(userInfoUrl)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	info := &OAUserInfo{}
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}
	return info, err
}
