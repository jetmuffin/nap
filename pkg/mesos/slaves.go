package mesos

import (
	"encoding/json"
	"fmt"
)

// SlavesWrap is the wrapper for slave endpoint of mesos
type SlavesWrap struct {
	List []Slave `json:"slaves"`
}

// ListSlaves returns list of slaves.
func (c *Client) ListSlaves() ([]Slave, error) {
	body, err := c.doRequest(c.slavesEndpoint())
	if err != nil {
		return []Slave{}, err
	}

	var slaves SlavesWrap
	err = json.Unmarshal(body, &slaves)
	if err != nil {
		return []Slave{}, err
	}

	return slaves.List, nil
}

// GetSlaveByID returns specific slave with given ID.
func (c *Client) GetSlaveByID(ID string) (Slave, error) {
	slaves, err := c.ListSlaves()
	if err != nil {
		return Slave{}, err
	}

	for _, slave := range slaves {
		if slave.ID == ID {
			return slave, nil
		}
	}

	return Slave{}, fmt.Errorf("slave ID %s not found", ID)
}
