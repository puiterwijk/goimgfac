package goimgfac

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type imgFacClient struct {
	baseurl      string
	clientid     string
	clientsecret string
}

func CreateImgFacClient(baseurl string, clientid string, clientsecret string) *imgFacClient {
	return &imgFacClient{baseurl, clientid, clientsecret}
}

func (c *imgFacClient) call(method string,
	urlParts []string,
	arguments *map[string]string) ([]byte, error) {
	u, _ := url.Parse(c.baseurl)
	u.Path = "/imagefactory"

	if len(urlParts) != 0 {
		u.Path += "/"
	}
	u.Path += strings.Join(urlParts, "/")
	q := url.Values{}
	if arguments != nil {
		for k, v := range *arguments {
			if v != "" {
				q.Add(k, v)
			}
		}
	}

	if method != "POST" {
		u.RawQuery = q.Encode()
	}

	var req *http.Request
	var err error
	if method == "POST" {
		req, err = http.NewRequest("POST", u.String(), bytes.NewReader([]byte(q.Encode())))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(method, u.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
