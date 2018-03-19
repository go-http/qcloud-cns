package cns

import (
	"net/url"
	"strconv"
)

//“域名”的数据结构
type Domain struct {
	Id               int
	Status           string
	GroupId          string `json:"group_id"`
	SearchenginePush string `json:"searchengine_push"`
	IsMark           string `json:"is_mark"`
	Ttl              string
	CnameSpeedup     string `json:"cname_speedup"`
	Remark           string
	CreatedOn        string `json:"created_on"`
	UpdatedOn        string `json:"updated_on"`
	QProjectId       int    `json:"q_project_id"`
	Punycode         string
	ExtStatus        string `json:"ext_status"`
	SrcFlag          string `json:"src_flag"`
	Name             string
	Grade            string `json:"DP_Free"`
	GradeTitle       string `json:"grade_title"`
	IsVip            string `json:"is_vip"`
	Owner            string
	Records          string
	MinTtl           int `json:"min_ttl"`
}

//获取域名列表
func (cli *Client) DomainList() ([]Domain, error) {
	var respInfo struct {
		BaseResponse
		Data struct {
			Info struct {
				DomainTotal int `json:"domain_total"`
			}
			Domains []Domain
		}
	}
	err := cli.requestGET("DomainList", nil, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Data.Domains, nil
}

//添加域名，如果成功，返回创建的域名ID
func (cli *Client) DomainCreate(domain string, projectId ...int) (int, error) {
	param := url.Values{"domain": {domain}}

	if len(projectId) > 0 {
		param.Set("projectId", strconv.Itoa(projectId[0]))
	}

	var respInfo struct {
		BaseResponse
		Data struct {
			Domain struct {
				Id       int `json:",string"`
				Punycode string
				Domain   string
			}
		}
	}

	err := cli.requestGET("DomainCreate", param, &respInfo)
	if err != nil {
		return 0, err
	}

	return respInfo.Data.Domain.Id, nil
}
