package goimgfac

import (
	"encoding/json"
)

type serverInfoResult struct {
	ApiVersion string `json:"api_version"`
	Name       string `json:"name"`
	Version    string `json:"version"`
}

func (c *imgFacClient) ServerInfo() (*serverInfoResult, error) {
	resp, err := c.call("GET", []string{}, nil)
	if err != nil {
		return nil, err
	}
	var result serverInfoResult
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
