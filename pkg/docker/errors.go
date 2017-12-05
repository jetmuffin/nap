package docker

import (
	"fmt"
)

// errConnectionFailed implements an error returned when connection failed.
type errConnectionFailed struct {
	host string
}

// Error returns a string representation of an errConnectionFailed
func (err errConnectionFailed) Error() string {
	if err.host == "" {
		return "Cannot connect to the Docker daemon. Is the docker daemon running on this host?"
	}
	return fmt.Sprintf("Cannot connect to the Docker daemon at %s. Is the docker daemon running?", err.host)
}

// ErrorConnectionFailed returns an error with host in the error message when connection to docker daemon failed.
func ErrorConnectionFailed(host string) error {
	return errConnectionFailed{host: host}
}
