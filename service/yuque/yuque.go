package yuque

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

var appName = "bastion"
var repoID = ""
var token = ""

var client *resty.Client

func init() {
	client = resty.New()
}

func GetAllDocs() ([]*Doc, error) {
	res := DocsRes{}
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Auth-Token", token).
		SetHeader("User-Agent", appName).
		SetResult(&res).
		Get(fmt.Sprintf("https://www.yuque.com/api/v2/repos/%v/docs", repoID))

	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func GetDocDetail(id string) (*DocRes, error) {
	res := DocRes{}
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Auth-Token", token).
		SetHeader("User-Agent", appName).
		SetResult(&res).
		Get(fmt.Sprintf("https://www.yuque.com/api/v2/repos/%v/docs/%v", repoID, id))

	if err != nil {
		return nil, err
	}
	return &res, nil
}
