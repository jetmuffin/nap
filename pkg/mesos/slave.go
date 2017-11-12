package mesos

import (
	"fmt"
	. "github.com/JetMuffin/nap/pkg/types"
)

// GetSlaveByID will return a a slave by its unique ID (slaveID).
//
// The list of slaves are provided by a state of a single node.
func (c *Client) GetSlaveByID(slaves []MesosSlave, slaveID string) (*MesosSlave, error) {
	for _, slave := range slaves {
		if slaveID == slave.ID {
			return &slave, nil
		}
	}

	return nil, fmt.Errorf("No slave found with id \"%s\"", slaveID)
}

