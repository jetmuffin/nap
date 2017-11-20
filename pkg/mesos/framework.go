package mesos

import (
	"fmt"
	. "github.com/JetMuffin/nap/pkg/types"
	"strings"
)

// GetFrameworkByPrefix will return a framework that matches prefix.
//
// The list of framework are provided by a state of a slave / master.
func (c *Client) GetFrameworkByPrefix(frameworks []MesosFramework, prefix string) (*MesosFramework, error) {
	for _, framework := range frameworks {
		if strings.HasPrefix(framework.Name, prefix) {
			return &framework, nil
		}
	}

	return nil, fmt.Errorf("Framework with prefix \"%s\" not found", prefix)
}
