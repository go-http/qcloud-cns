package cns

import (
	"net/url"
	"testing"
)

func TestSignature(t *testing.T) {
	method := "GET"
	addr := "cvm.api.qcloud.com/v2/index.php"
	secretKey := "Gu5t9xGARNpq86cd98joQYCN3Cozk1qA"
	param := url.Values{
		"Action":        {"DescribeInstances"},
		"InstanceIds.0": {"ins-09dx96dg"},
		"Nonce":         {"11886"},
		"Region":        {"ap-guangzhou"},
		"SecretId":      {"AKIDz8krbsJ5yKBZQpn74WFkmLPx3gnPhESA"},
		"Timestamp":     {"1465185768"},
	}

	sigSha1 := "nPVnY6njQmwQ8ciqbPl5Qe+Oru4="
	sigSha256 := "0EEm/HtGRr/VJXTAD9tYMth1Bzm3lLHz5RCDv1GdM8s="

	var sig string

	param.Set("SignatureMethod", "HmacSHA1")
	sig = Signature(param, method, addr, secretKey)
	if sig != sigSha1 {
		t.Errorf("加密错误: %s != %s", sig, sigSha1)
	}

	param.Set("SignatureMethod", "HmacSHA256")
	sig = SignatureSha256(param, method, addr, secretKey)
	if sig != sigSha256 {
		t.Errorf("加密错误: %s != %s", sig, sigSha256)
	}
}
