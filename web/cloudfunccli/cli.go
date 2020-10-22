package cloudfunccli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	AccessToken = "38_hUOAM3ML1uFeMomS_3Vlm9uxzaOLvr9JGeJhyLT5PTJ3l6XFe7vqUkLBc1EJwQP0Ix4wYV19rvIRyy37XPyYJoaco47lAn1RSNk-1u_1FcFHRyrEaZsTnnA-m8DCqaOmujEloD9EjXyZ4z19BOZjACATGR"
	OATotoken   = ""
)

const (
	// env = "dev-osmu3"
	env     = "release-8tcge"
	baseURL = "https://api.weixin.qq.com/tcb/invokecloudfunction"
)

// DoRequest requests the mini progrom cloudfunctions
func DoRequest(funcName string, params string) (string, error) {
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, AccessToken, env, funcName), "application/json", strings.NewReader(params))
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
