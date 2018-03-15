package cns

import (
	"net/url"
)

//“域名”“解析记录”的数据结构
type Record struct {
	Id         int
	Ttl        int
	Value      string
	Enabled    int
	Status     string
	UpdatedOn  string `json:"updated_on"`
	QProjectId int    `json:"q_project_id"`
	Name       string
	Line       string
	LineId     string `json:"line_id"`
	Type       string
	Remark     string
	Mx         int
	Hold       string
}

//“域名”“解析记录”列表中的“域名”数据结构，和“域名列表”中的略有不同
type DomainInRecordList struct {
	Domain
	Ttl      int
	Id       int      `json:",string"`
	DnspodNs []string `json:"dnspod_ns"`
}

//获取指定“域名”的“解析记录”列表
func (cli *Client) RecordList(domain string) ([]Record, error) {
	var respInfo struct {
		BaseResponse
		Data struct {
			Domain  DomainInRecordList
			Records []Record
			Info    struct {
				SubDomains  int `json:"sub_domains,string"`
				RecordTotal int `json:"record_total,string"`
			}
		}
	}

	param := url.Values{
		"domain": {domain},
	}

	err := cli.requestGET("RecordList", param, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Data.Records, nil
}
