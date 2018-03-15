package cns

//云解析API请求的客户端
type Client struct {
	SecretId  string
	SecretKey string
}

func New(secretId, secretKey string) *Client {
	return &Client{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}
