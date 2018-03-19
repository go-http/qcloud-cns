package cns

import (
	"net/url"
	"strconv"
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

//添加指定“域名”的“解析记录”
func (cli *Client) RecordCreate(domain string, record Record) (int, error) {
	if record.Line == "" {
		record.Line = "默认"
	}

	//必选参数
	param := url.Values{
		"domain":     {domain},
		"subDomain":  {record.Name},
		"recordType": {record.Type},
		"recordLine": {record.Line},
		"value":      {record.Value},
	}

	//可选TTL参数，缺省为600
	if record.Ttl > 0 {
		param.Set("ttl", strconv.Itoa(record.Ttl))
	}

	//MX记录必须的额外参数
	if record.Type == "MX" {
		param.Set("mx", strconv.Itoa(record.Mx))
	}

	var respInfo struct {
		BaseResponse
		Data struct {
			Record struct {
				Id     int `json:",string"`
				Name   string
				Status string
				Weight interface{}
			}
		}
	}
	err := cli.requestGET("RecordCreate", param, &respInfo)
	if err != nil {
		return 0, err
	}

	return respInfo.Data.Record.Id, nil
}

//设置指定“域名”的“解析记录”状态
func (cli *Client) RecordStatus(domain string, recordId int, enable bool) error {
	param := url.Values{
		"domain":   {domain},
		"recordId": {strconv.Itoa(recordId)},
	}

	if enable {
		param.Set("status", "enable")
	} else {
		param.Set("status", "disable")
	}

	var respInfo BaseResponse
	err := cli.requestGET("RecordStatus", param, &respInfo)
	if err != nil {
		return err
	}

	return nil
}
