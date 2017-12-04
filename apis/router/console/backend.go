package console

import (
	"github.com/JetMuffin/nap/pkg/mesos"
)

// Backend abstracts a console manager
type Backend interface {
	GetTaskByID(taskID string) (mesos.Task, error)
	GetTaskContainerName(task mesos.Task) (string, error)
	GetSlaveByID(slaveID string) (mesos.Slave, error)
}
