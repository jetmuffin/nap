package mesos

import (
	"encoding/json"
	"fmt"
)

// FrameworksWrap is the wrapper for mesos framework endpoints.
type FrameworksWrap struct {
	List []Framework `json:"frameworks"`
}

// ListFrameworks returns list of current running frameworks.
func (c *Client) ListFrameworks() ([]Framework, error) {
	body, err := c.doRequest(c.frameworksEndpoint())
	if err != nil {
		return []Framework{}, err
	}

	var frameworks FrameworksWrap
	err = json.Unmarshal(body, &frameworks)
	if err != nil {
		return []Framework{}, err
	}

	return frameworks.List, nil
}

// GetFrameworkByID returns specific framework with given id.
func (c *Client) GetFrameworkByID(ID string) (Framework, error) {
	frameworks, err := c.ListFrameworks()
	if err != nil {
		return Framework{}, err
	}

	for _, framework := range frameworks {
		if framework.ID == ID {
			return framework, nil
		}
	}

	return Framework{}, fmt.Errorf("framework ID %s not found", ID)
}
