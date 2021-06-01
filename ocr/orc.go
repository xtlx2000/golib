package ocr

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	goliblog "github.com/xtlx2000/golib/log"
)

const (
	// 客户端凭证类型，固定为client_credentials
	grantType string = "client_credentials"
	// 应用的API Key
	apiKey string = "hWa2LhCzAwYSmhG3th4HDlDK"
	// 应用的Secret Key
	secretKey string = "yGRZjm4433xLzZYCyXrgG8Pyzfzli9Sd"
	// 授权服务地址
	tokenUrl string = "https://aip.baidubce.com/oauth/2.0/token"
	// 文字识别（高精度 500次/天）API接口地址
	// accurateBasicUrl string = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
	// 文字识别（标准版 50000次/天）API接口地址
	accurateBasicUrl string = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
)

type accessToken struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     uint32 `json:"expires_in"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
}

type WordsResult struct {
	Words string `json:"words"`
}

type Words struct {
	WordsResult    []WordsResult `json:"words_result"`
	LogId          uint64        `json:"log_id"`
	WordsResultNum uint32        `json:"words_result_num"`
}

// 获取access_token
// access_token有效期一般是30天
// 官方文档：https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu
func getAccessToken() (data accessToken, err error) {
	requestUrl := fmt.Sprintf("%s?grant_type=%s&client_id=%s&client_secret=%s", tokenUrl, grantType, apiKey, secretKey)
	response, err := http.Get(requestUrl)
	if err != nil {
		goliblog.Errorf("http get error: %v", err)
		return
	}
	defer response.Body.Close()
	s := unCoding(response)
	err = json.Unmarshal([]byte(s), &data)
	if err != nil {
		goliblog.Errorf("unmarshal error:%v", err)
	}
	return
}

// 文字识别（高精度）
// 官方文档：https://cloud.baidu.com/doc/OCR/s/1k3h7y3db
// img 图片地址
func Recognite(img string) (data Words, err error) {
	token, err := getAccessToken()
	if err != nil {
		goliblog.Errorf("getAccessToken error: %v", err)
		return
	}
	requestUrl := fmt.Sprintf("%s?access_token=%s", accurateBasicUrl, token.AccessToken)
	f, e := os.Open(img)
	printError(e)
	defer f.Close()
	d, e := ioutil.ReadAll(f)
	printError(e)
	const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var coder = base64.NewEncoding(base64Table)
	imgString := coder.EncodeToString(d)
	printError(e)
	values := url.Values{"image": {imgString}}
	response, e := http.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	defer response.Body.Close()
	s := unCoding(response)
	e = json.Unmarshal([]byte(s), &data)
	printError(e)
	return
}

func unCoding(r *http.Response) (body string) {
	if r.StatusCode == 200 {
		switch r.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(r.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				body += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(r.Body)
			body = string(bodyByte)
		}
	} else {
		bodyByte, _ := ioutil.ReadAll(r.Body)
		body = string(bodyByte)
	}
	return
}

func printError(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}

/*
func main() {
	fmt.Println(Recognite("./test.jpg").WordsResult[0].Words)
}
*/
