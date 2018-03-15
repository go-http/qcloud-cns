package cns

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

//使用SHA1签名请求
func Signature(param url.Values, reqMethod, addr, secretKey string) string {
	return signature(param, reqMethod, addr, secretKey, "sha1")
}

//使用SHA256签名请求
func SignatureSha256(param url.Values, reqMethod, addr, secretKey string) string {
	return signature(param, reqMethod, addr, secretKey, "sha256")
}

//按指定签名方法签名请求
func signature(param url.Values, reqMethod, addr, secretKey, sigMethod string) string {
	strs := make([]string, 0, len(param))
	for key, values := range param {
		var value string
		if len(values) == 1 {
			value = values[0]
		}

		key = strings.Replace(key, "_", ".", -1)

		strs = append(strs, key+"="+value)
	}

	sort.Strings(strs)

	str := strings.Join(strs, "&")

	data := []byte(reqMethod + addr + "?" + str)

	//指定sha256则用sha256，否则一律使用sha1
	hashFunc := sha1.New
	if sigMethod == "sha256" {
		hashFunc = sha256.New
	}

	mac := hmac.New(hashFunc, []byte(secretKey))
	mac.Write([]byte(data))
	sigData := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(sigData)
}
