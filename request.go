package cns

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

//API响应的接口
type Responser interface {
	Error() error
}

//基础的API响应结构
type BaseResponse struct {
	Code     int
	Message  string
	CodeDesc string
}

//API响应中识别并构造API错误的方法
func (resp BaseResponse) Error() error {
	if resp.Code == 0 {
		return nil
	}

	return fmt.Errorf("[%d](%s): %s", resp.Code, resp.CodeDesc, resp.Message)
}

//云解析API请求的资源地址
const Uri = "cns.api.qcloud.com/v2/index.php"

//GET类型的API请求封装
func (cli *Client) requestGET(action string, param url.Values, respInfo Responser) error {
	return cli.request("GET", action, param, nil, respInfo)
}

//API请求的封装（内建公共参数、签名的设置）
func (cli *Client) request(method, action string, param url.Values, body io.Reader, respInfo Responser) error {
	if param == nil {
		param = url.Values{}
	}

	//设置公共参数
	param.Set("Action", action)
	param.Set("Timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	param.Set("Nonce", "123456")
	param.Set("SecretId", cli.SecretId)

	sig := Signature(param, method, Uri, cli.SecretKey)
	param.Set("Signature", sig)

	req, err := http.NewRequest(method, "https://"+Uri+"?"+param.Encode(), body)
	if err != nil {
		return fmt.Errorf("构建请求错误: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("执行请求错误: %s", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(respInfo)
	if err != nil {
		return fmt.Errorf("读取响应错误: %s", err)
	}

	return respInfo.Error()
}
