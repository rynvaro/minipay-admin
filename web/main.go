package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

var access_token = "38_8_BwGGTv69yaep-yK9UnClnZgYwDkzteHhYC6Snod5gIaWXj7mz7oW7GTsv_4Xn_Nh04bDrcZZVNX3d7Z10fO4tilX530ZMpzoEiNi6en1FDqiwr53vD64GJlLpmdOjIXd9H5RiHNQfuV1I8AXKhAIAIFS"

const (
	env = "dev-osmu3"
	// env      = "release-8tcge"
	baseURL  = "https://api.weixin.qq.com/tcb/invokecloudfunction"
	funcName = "console"
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
			access_token = r.AccessToken
			fmt.Println("token is: ", access_token)
			time.Sleep(time.Minute * 60)
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
	r.POST("/stores", stores)
	r.GET("/plates", plates)
	r.POST("/updateplate", updateplate)
	r.POST("/addplate", addplate)
	r.GET("/plate", plate)
	r.GET("/users", users)
	r.POST("/updateuser", updateuser)
	r.POST("/addcoupon", addcoupon)
	r.GET("/store", store)
	r.POST("/storedelete", storedelete)

	r.GET("/events", events)
	r.POST("/eventadd", eventadd)
	r.POST("/eventdelete", eventdelete)
	r.POST("/eventupdate", eventupdate)

	r.GET("/batch", batch)

	r.POST("/proxy", proxy)

	r.OPTIONS("/upload", options)
	r.OPTIONS("/publish", options)
	r.OPTIONS("/stores", options)
	r.OPTIONS("/plates", options)
	r.OPTIONS("/updateplate", options)
	r.OPTIONS("/users", options)
	r.OPTIONS("/updateuser", options)
	r.OPTIONS("/addcoupon", options)
	r.OPTIONS("/store", options)
	r.OPTIONS("/addplate", options)
	r.OPTIONS("/plate", options)
	r.OPTIONS("/events", options)
	r.OPTIONS("/eventadd", options)
	r.OPTIONS("/eventdelete", options)
	r.OPTIONS("/eventupdate", options)
	r.OPTIONS("/storedelete", options)
	r.OPTIONS("/proxy", options)
	r.OPTIONS("/batch", options)

	box := packr.NewBox("dist")
	static := packr.NewBox("dist/static")
	r.StaticFS("/web", box)
	r.StaticFS("static", static)

	go func() {
		time.Sleep(time.Second)
		indexUrl := "http://localhost:8090/web/#/components/pubstores"
		switch runtime.GOOS {
		case "windows":
			exec.Command(`cmd`, `/c`, `start`, indexUrl).Start()
		case "darwin":
			exec.Command(`open`, indexUrl).Start()
		case "linux":
			exec.Command(`xdg-open`, indexUrl).Start()
		}
	}()

	r.Run(":8090")
}

type Store struct {
	Id               string    `json:"_id"`
	Address          string    `json:"address"`
	AvgPrice         float64   `json:"avgPrice"`
	Balance          float64   `json:"balance"`
	Bank             string    `json:"bank"`
	Banners          []*Banner `json:"banners"`
	CreatedAt        int64     `json:"createdAt"`
	Deleted          int       `json:"deleted"`
	Discount         float64   `json:"discount"`
	DiscountStart    int64     `json:"discountStart"`
	DiscountEnd      int64     `json:"discountEnd"`
	EndTime          string    `json:"endTime"`
	GeoPoint         *GeoPoint `json:"geoPoint"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	MerchantBankCard string    `json:"merchantBankCard"`
	MerchantName     string    `json:"merchantName"`
	MerchantPhone    string    `json:"merchantPhone"`
	OpenID           string    `json:"openid"`
	Orders           int       `json:"orders"`
	Password         string    `json:"password"`
	ProductImages    []string  `json:"productImages"`
	PromoteImages    []string  `json:"promoteImages"`
	StartTime        string    `json:"startTime"`
	StoreDesc        string    `json:"storeDesc"`
	StoreImages      []string  `json:"storeImages"`
	StoreName        string    `json:"storeName"`
	StoreType        int       `json:"storeType"`
	UpdatedAt        int64     `json:"updatedAt"`
	UserInfo         *UserInfo `json:"userInfo"`
	TempFileURL      string    `json:"tempFileUrl"`
}

type UserInfo struct {
	AppID string `json:"appId"`
}

type GeoPoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Stores struct {
	Data []*Store `json:"data"`
}

func batch(c *gin.Context) {
	params := `{"currentPage":1, "tp":"list"}`
	resp, err := doRequest("aaaaaaaa", params)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	stores := &Stores{}
	if err := json.Unmarshal([]byte(resp), stores); err != nil {
		panic(err)
	}

	fileIds := make([]string, 0)
	for _, v := range stores.Data {
		if len(v.PromoteImages) > 0 {
			fileIds = append(fileIds, v.PromoteImages[0])
		}

	}
	respFile, err := getTempFiles(fileIds)
	if err != nil {
		panic(err)
	}
	respFileMap := map[string]string{}
	for _, v := range respFile {
		respFileMap[v.FileID] = v.DownloadURL
	}

	fmt.Println(len(stores.Data))
	for i := range stores.Data {
		if len(stores.Data[i].PromoteImages) > 0 {
			if durl, ok := respFileMap[stores.Data[i].PromoteImages[0]]; ok {
				fmt.Println("resizing...", stores.Data[i].StoreName)
				stores.Data[i].TempFileURL = durl
				fileResp, err := http.Get(durl)
				if err != nil {
					panic(err)
				}
				downfileName := fmt.Sprintf("%v%v", time.Now().UnixNano(), filepath.Ext(durl))
				out, err := os.Create(downfileName)
				if err != nil {
					panic(err)
				}
				_, err = io.Copy(out, fileResp.Body)
				if err != nil {
					panic(err)
				}
				out.Close()

				file, err := os.Open(downfileName)
				if err != nil {
					log.Fatal(err)
				}

				var img image.Image
				if filepath.Ext(downfileName) == ".png" {
					img, err = png.Decode(file)
				} else {
					img, err = jpeg.Decode(file)
				}
				if err != nil {
					log.Fatal(err)
				}
				file.Close()
				os.Remove(file.Name())

				// resize to width 1000 using Lanczos resampling
				// and preserve aspect ratio
				m := resize.Resize(340, 0, img, resize.Lanczos3)

				genFile := "gen" + downfileName
				out, err = os.Create(genFile)
				if err != nil {
					log.Fatal(err)
				}

				// write new image to file
				if filepath.Ext(downfileName) == ".png" {
					err = png.Encode(out, m)
				} else {
					err = jpeg.Encode(out, m, nil)
				}
				out.Close()
				if err != nil {
					fmt.Println(err)
				}
				// upload image
				fid := doUpload(genFile)

				params := `{"storeId":"` + stores.Data[i].Id + `", "tp":"update", "thumbnail":"` + fid + `"}`
				_, err = doRequest("aaaaaaaa", params)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("id: ", stores.Data[i].Id)
			}
		}
	}

	c.JSON(http.StatusOK, stores)
}

func proxy(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp, err := doRequest("consolecommon", string(data))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusOK, resp)
}

func storedelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resp, err := doRequest("consolecommon", string(data))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusOK, resp)
}

func eventdelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resp, err := doRequest("consolecommon", string(data))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusOK, resp)
}

func eventupdate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resp, err := doRequest("consolecommon", string(data))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusOK, resp)
}

func eventadd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resp, err := doRequest("consolecommon", string(data))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusOK, resp)
}

func events(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := `{"tp":"events"}`

	resp, err := doRequest("consolecommon", params)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusOK, resp)
}

func plate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id := c.Request.FormValue("id")
	params := `{"_id":"` + id + `", "tp": 3}`

	resp, err := doRequest("consoleplates", params)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusOK, resp)
}

func doRequest(funcName string, params string) (string, error) {
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, funcName), "application/json", strings.NewReader(params))
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// fmt.Println(string(data))
	r := &Resp{}
	if err := json.Unmarshal(data, r); err != nil {
		return "", err
	}
	if r.Errcode != 0 {
		return "", fmt.Errorf("faied")
	}
	return r.RespData, nil
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

	fileList, err := getTempFiles(images)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	rf := &struct {
		FileList []*RespFileItem `json:"file_list"`
		RespData string          `json:"resp_data"`
	}{
		FileList: fileList,
	}
	rf.RespData = r.RespData

	c.JSON(http.StatusOK, rf)
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

/**
{"env":"dev-osmu3","file_list":[{"fileid":"cloud://dev-osmu3.6465-dev-osmu3-1302781643/images/console/f0a99207-8a44-4e99-bae8-22ffa05382e7.mp4","max_age":7200},
{"fileid":"cloud://dev-osmu3.6465-dev-osmu3-1302781643/images/console/7d017864-03ba-474c-ae68-18d2f693d6d3.jpg","max_age":7200},
{"fileid":"cloud://dev-osmu3.6465-dev-osmu3-1302781643/images/console/14460677-1be8-4a71-bce7-6dcff2c1efd0.jpg","max_age":7200},
{"fileid":"cloud://dev-osmu3.6465-dev-osmu3-1302781643/images/console/c80715bf-f134-4d9a-9ed5-e2b815a34b3d.jpg","max_age":7200}]}

{"errcode":0,"errmsg":"ok","file_list":[
	{"fileid":"cloud:\/\/dev-osmu3.6465-dev-osmu3-1302781643\/images\/console\/f0a99207-8a44-4e99-bae8-22ffa05382e7.mp4",
	"download_url":"https:\/\/6465-dev-osmu3-1302781643.tcb.qcloud.la\/images\/console\/f0a99207-8a44-4e99-bae8-22ffa05382e7.mp4",
	"status":0,"errmsg":"ok"},
	{"fileid":"cloud:\/\/dev-osmu3.6465-dev-osmu3-1302781643\/images\/console\/7d017864-03ba-474c-ae68-18d2f693d6d3.jpg",
	"download_url":"https:\/\/6465-dev-osmu3-1302781643.tcb.qcloud.la\/images\/console\/7d017864-03ba-474c-ae68-18d2f693d6d3.jpg","status":0,"errmsg":"ok"},
	{"fileid":"cloud:\/\/dev-osmu3.6465-dev-osmu3-1302781643\/images\/console\/14460677-1be8-4a71-bce7-6dcff2c1efd0.jpg",
	"download_url":"https:\/\/6465-dev-osmu3-1302781643.tcb.qcloud.la\/images\/console\/14460677-1be8-4a71-bce7-6dcff2c1efd0.jpg","status":0,"errmsg":"ok"},
	{"fileid":"cloud:\/\/dev-osmu3.6465-dev-osmu3-1302781643\/images\/console\/c80715bf-f134-4d9a-9ed5-e2b815a34b3d.jpg",
	"download_url":"https:\/\/6465-dev-osmu3-1302781643.tcb.qcloud.la\/images\/console\/c80715bf-f134-4d9a-9ed5-e2b815a34b3d.jpg","status":0,"errmsg":"ok"}]}
*/

func getTempFiles(fileIds []string) (fileList []*RespFileItem, err error) {

	fileListReq := make([]*FileItem, 0)
	for i := range fileIds {
		item := &FileItem{
			FileID: fileIds[i],
			MaxAge: 7200,
		}
		fileListReq = append(fileListReq, item)
	}

	reqParams := struct {
		Env      string      `json:"env"`
		FileList []*FileItem `json:"file_list"`
	}{
		Env:      env,
		FileList: fileListReq,
	}

	data, err := json.Marshal(reqParams)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/tcb/batchdownloadfile?access_token=%s", access_token), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)

	rf := &struct {
		ErrCode  int             `json:"errcode"`
		ErrMsg   string          `json:"errmsg"`
		FileList []*RespFileItem `json:"file_list"`
	}{}

	if rf.ErrCode != 0 {
		return nil, fmt.Errorf(rf.ErrMsg)
	}

	if err = json.Unmarshal(data, rf); err != nil {
		return nil, err
	}
	return rf.FileList, nil
}

type RespFileItem struct {
	FileID      string `json:"fileid"`
	DownloadURL string `json:"download_url"`
}

type FileItem struct {
	FileID string `json:"fileid"`
	MaxAge int32  `json:"max_age"`
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

func addplate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println("req is: ", string(data))
	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consoleplates"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return

	}

	data, err = ioutil.ReadAll(resp.Body)
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

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", baseURL, access_token, env, "consolestores"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
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
	// fmt.Println(string(data))
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

func doUpload(fpath string) string {

	file, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}

	uuid, _ := uuid.NewRandom()
	path := "images/console/" + uuid.String() + filepath.Ext(file.Name())

	body := `
	{
		"env": "` + env + `",
		"path": "` + path + `"
	}`
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/tcb/uploadfile?access_token=%s", access_token), "application/json", strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(data))
	uR := &uploadResp{}
	if err := json.Unmarshal(data, uR); err != nil {
		panic(err)
	}

	if uR.Errcode != 0 {
		panic(err)
	}
	multipartBody := &bytes.Buffer{}
	writer := multipart.NewWriter(multipartBody)

	params := map[string]string{"key": path, "Signature": uR.Authorization, "x-cos-security-token": uR.Token, "x-cos-meta-fileid": uR.CosFileId}

	for key, val := range params {
		err = writer.WriteField(key, val)
		if err != nil {
			panic(err)
		}
	}

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		panic(err)
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if _, err := part.Write(fileData); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}

	resp, err = http.Post(uR.Url, writer.FormDataContentType(), multipartBody)
	if err != nil {
		panic(err)
	}
	// data, err = ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)
	os.Remove(fpath)

	return uR.FileID
}
