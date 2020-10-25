package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

const (
	token   = "38_1xNG7MF_-M2zy457GKa4YEoEIlWspQlDuJ3f_jAsWqqN1zg5566v_jkemM8KyeaYvmvAy4sjtW5D-2Ox9D9RCP3HPvg3A8yAFKs8jTTA7ZCNTkSQpemc8YQqEVBcmCE67k5Zp1C5kPz5Ci7iQDEcAGATGH"
	env     = "dev-osmu3"
	baseURL = "https://api.weixin.qq.com/tcb/invokecloudfunction"
)

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go DoRequest("aaaaaaaa", "{}")
	}
	wg.Wait()
}

// DoRequest requests the mini progrom cloudfunctions
func DoRequest(funcName string, params string) (string, error) {
	defer wg.Done()
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, token, env, funcName), "application/json", strings.NewReader(params))
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
