// Official Account
package oa

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"web/cloudfunccli"

	"github.com/gin-gonic/gin"
)

const oaToken = "thisistokenforqihaoshenghuo"

// 公众号回调认证
func Authentication(c *gin.Context) {
	signature := c.Request.FormValue("signature")
	t := c.Request.FormValue("timestamp")
	nonce := c.Request.FormValue("nonce")
	echoStr := c.Request.FormValue("echostr")

	fmt.Println("signature: ", signature)
	fmt.Println("timestamp: ", t)
	fmt.Println("nonce: ", nonce)
	fmt.Println("echoStr: ", echoStr)

	arr := []string{oaToken, t, nonce}
	sort.Strings(arr)
	fmt.Println("arr: ", arr)

	str := strings.Join(arr, "")
	fmt.Println("str: ", str)

	h := sha1.New()
	h.Write([]byte(str))
	cSignature := hex.EncodeToString(h.Sum(nil))
	fmt.Println("cSignature: ", cSignature)

	if cSignature == signature {
		c.String(http.StatusOK, echoStr)
		return
	}
	c.Status(http.StatusNotAcceptable)
}

func OACallback(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusNotAcceptable)
		return
	}

	fmt.Println(string(data))
	followEvent := &FollowEvent{}
	if err := xml.Unmarshal(data, followEvent); err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, "")
		return
	}
	oaUserInfo, err := GetOAUserInfo(cloudfunccli.OATotoken, followEvent.FromUserName)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, "")
		return
	}

	oaUserInfo.Tp = "oacallback"
	jsonData, err := json.Marshal(oaUserInfo)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, "")
		return
	}
	fmt.Println(string(jsonData))
	cloudfunccli.DoRequest("consolecommon", string(jsonData))

	c.String(http.StatusOK, "")
}

type FollowEvent struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
}
