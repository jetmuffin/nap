package mesos

import (
	"fmt"
	. "github.com/JetMuffin/nap/pkg/types"
)

// GetExecutorByID will return an Executor by its unique ID.
//
// The list of executors are provided by a framework.
func (c *Client) GetExecutorByID(executor []MesosExecutor, executorID string) (*MesosExecutor, error) {
	for _, e := range executor {
		if e.ID == executorID {
			return &e, nil
		}
	}

	return nil, fmt.Errorf("No executor found with id \"%s\"", executorID)
}
