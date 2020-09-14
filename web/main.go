package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var access_token = "37_D0B76czxEg51-r5VCL0SOVQl9Q61wIMfycScmbhJQ9bl0Y0YV1PHtBYBc-Tj4gbE_qIQF7ZWv0Yo49XU8Uj4zj6rAIi9jijgPCNOBPzTYoUTp2Gw7a4qsa6fxacLXY9l1LCqOV6gNo0O-pvjFZMeAGAKER"

const (
	env      = "dev-osmu3"
	baseURL  = "https://api.weixin.qq.com/tcb/invokecloudfunction"
	funcName = "console"
)

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
			access_token = r.AccessToken
			fmt.Println("token is: ", access_token)
			time.Sleep(time.Minute * 20)
		}
	}()
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
}

func main() {

	r := gin.Default()

	r.POST("/publish", publish)
	r.POST("/upload", upload)
	r.GET("/stores", stores)
	r.GET("/plates", plates)
	r.POST("/updateplate", updateplate)
	r.GET("/users", users)
	r.POST("/updateuser", updateuser)
	r.POST("/addcoupon", addcoupon)
	r.GET("/store", store)

	r.OPTIONS("/upload", options)
	r.OPTIONS("/publish", options)
	r.OPTIONS("/stores", options)
	r.OPTIONS("/plates", options)
	r.OPTIONS("/updateplate", options)
	r.OPTIONS("/users", options)
	r.OPTIONS("/updateuser", options)
	r.OPTIONS("/addcoupon", options)
	r.OPTIONS("/store", options)

	// box := packr.NewBox("dist")
	// static := packr.NewBox("dist/static")
	// r.StaticFS("/web", box)
	// r.StaticFS("static", static)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	indexUrl := "http://localhost:8090/web/#/components/pubstores"
	// 	switch runtime.GOOS {
	// 	case "windows":
	// 		exec.Command(`cmd`, `/c`, `start`, indexUrl).Start()
	// 	case "darwin":
	// 		exec.Command(`open`, indexUrl).Start()
	// 	case "linux":
	// 		exec.Command(`xdg-open`, indexUrl).Start()
	// 	}
	// }()

	r.Run(":8090")
}

func store(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id := c.Request.FormValue("id")
	params := `{"id":"` + id + `", "tp": 1}`
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consolestores"), "application/json", strings.NewReader(params))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	r := &Resp{}
	if err := json.Unmarshal(data, r); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if r.Errcode != 0 {
		c.Status(http.StatusInternalServerError)
		return
	}

	rd := &RespData{}
	if err := json.Unmarshal([]byte(r.RespData), rd); err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}
	images := make([]string, 0)

	for _, v := range rd.Data.Banners {
		images = append(images, v.URL)
	}

	images = append(images, rd.Data.StoreImages...)
	images = append(images, rd.Data.ProductImages...)
	images = append(images, rd.Data.PromoteImages...)

	fileList := make([]*FileItem, 0)
	for i := range images {
		item := &FileItem{
			FileID: images[i],
			MaxAge: 7200,
		}
		fileList = append(fileList, item)
	}

	reqParams := struct {
		Env      string      `json:"env"`
		FileList []*FileItem `json:"file_list"`
	}{
		Env:      env,
		FileList: fileList,
	}

	data, err = json.Marshal(reqParams)
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println(string(data))
	fmt.Println()

	resp, err = http.Post(fmt.Sprintf("https://api.weixin.qq.com/tcb/batchdownloadfile?access_token=%s", access_token), "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)

	rf := &RespFile{}

	if err := json.Unmarshal(data, rf); err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	rf.RespData = r.RespData

	c.JSON(http.StatusOK, rf)
}

type RespFile struct {
	FileList []*RespFileItem `json:"file_list"`
	RespData string          `json:"resp_data"`
}

type RespFileItem struct {
	FileID      string `json:"fileid"`
	DownloadURL string `json:"download_url"`
}

type FileItem struct {
	FileID string `json:"fileid"`
	MaxAge int32  `json:"max_age"`
}

type RespData struct {
	Data struct {
		Banners       []*Banner `json:"banners"`
		StoreImages   []string  `json:"storeImages"`
		ProductImages []string  `json:"productImages"`
		PromoteImages []string  `json:"promotoImages"`
	} `json:"data"`
}

type Banner struct {
	IsVideo bool   `json:"isVideo"`
	URL     string `json:"url"`
}

func addcoupon(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consoleupdateuser"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return

	}

	data, err = ioutil.ReadAll(resp.Body)

	fmt.Println(string(data), err)

	c.Status(http.StatusOK)
}

func updateuser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consoleupdateuser"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return

	}

	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)

	c.Status(http.StatusOK)
}

func users(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consoleusers"), "application/json", nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(data), err)
	r := &Resp{}
	if err := json.Unmarshal(data, r); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if r.Errcode != 0 {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, r.RespData)
}

func updateplate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consoleplates"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return

	}

	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)

	c.Status(http.StatusOK)
}

func plates(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consoleplates"), "application/json", nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(data))
	r := &Resp{}
	if err := json.Unmarshal(data, r); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if r.Errcode != 0 {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, r.RespData)
}

func stores(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	q := c.Request.FormValue("q")
	params := `{"q":"` + q + `"}`

	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consolestores"), "application/json", strings.NewReader(params))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(data))
	r := &Resp{}
	if err := json.Unmarshal(data, r); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if r.Errcode != 0 {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, r.RespData)
}

type Resp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errms"`
	RespData string `json:"resp_data"`
}

/**
200 OK
{"errcode":42001,"errmsg":"access_token expired rid: 5f5cafe9-456a174e-4e992ccd"}

{"errcode":0,"errmsg":"ok","resp_data":"{\"event\":{\"userInfo\":{\"appId\":\"wx9b588b2b3f090400\"}},\"appid\":\"wx9b588b2b3f090400\"}"}
*/
func publish(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, funcName), "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	c.Status(http.StatusOK)
}

func upload(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	uuid, _ := uuid.NewRandom()
	path := "images/console/" + uuid.String() + filepath.Ext(header.Filename)

	body := `
	{
		"env": "` + env + `",
		"path": "` + path + `"
	}`
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/tcb/uploadfile?access_token=%s", access_token), "application/json", strings.NewReader(body))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(data))
	uR := &uploadResp{}
	if err := json.Unmarshal(data, uR); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if uR.Errcode != 0 {
		c.Status(http.StatusInternalServerError)
		return
	}
	multipartBody := &bytes.Buffer{}
	writer := multipart.NewWriter(multipartBody)

	params := map[string]string{"key": path, "Signature": uR.Authorization, "x-cos-security-token": uR.Token, "x-cos-meta-fileid": uR.CosFileId}

	for key, val := range params {
		err = writer.WriteField(key, val)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := part.Write(fileData); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := writer.Close(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resp, err = http.Post(uR.Url, writer.FormDataContentType(), multipartBody)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)

	c.String(http.StatusOK, uR.FileID)
}

type uploadResp struct {
	Errcode       int    `json:"errcode"`
	Errmsg        string `json:"errms"`
	Url           string `json:"url"`
	Token         string `json:"token"`
	Authorization string `json:"authorization"`
	FileID        string `json:"file_id"`
	CosFileId     string `json:"cos_file_id"`
}

func options(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
	c.Status(http.StatusOK)
}
