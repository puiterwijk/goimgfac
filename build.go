package goimgfac

import (
	"encoding/json"
	"time"

	"fmt"
)

type baseImageBuildInfo struct {
	Type             string            `json:"_type"`
	Href             string            `json:"href"`
	Icicle           string            `json:"icicle"`
	Id               string            `json:"id"`
	Parameters       map[string]string `json:"parameters"`
	Percent_Complete int               `json:"percent_complete"`
	Status           string            `json:"status"`
	Status_Detail    map[string]string `json:"status_detail"`
	Template         string            `json:"template"`
}

type imageInfo struct {
	Base_Image baseImageBuildInfo `json:"base_image"`
}

func (c *imgFacClient) BuildBaseImage(template string,
	parameters map[string]string) (*imageInfo, error) {
	parameters["template"] = template
	fmt.Println("Parameters:", parameters)
	resp, err := c.call("POST",
		[]string{"base_images"},
		&parameters)
	if err != nil {
		return nil, err
	}
	var result imageInfo
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *imgFacClient) GetBaseImage(id string) (*imageInfo, error) {
	resp, err := c.call("GET", []string{"base_images", id}, nil)
	if err != nil {
		return nil, err
	}
	var result imageInfo
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *imgFacClient) WaitForBaseImageBuild(id string) error {
	var done bool = false
	for !done {
		info, err := c.GetBaseImage(id)
		if err != nil {
			return err
		}
		if info.Base_Image.Status != "NEW" && info.Base_Image.Status != "BUILDING" {
			done = true
		} else {
			time.Sleep(10 * time.Second)
		}
	}
	return nil
}
