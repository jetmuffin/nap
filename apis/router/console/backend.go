package console

import (
	mesos "github.com/JetMuffin/nap/pkg/types"
)

// Backend abstracts a console manager
type Backend interface {
	GetTaskByID(taskID string) (mesos.MesosTask, error)
	TaskContainerName(task mesos.MesosTask) (string, error)
	TaskSlave(task mesos.MesosTask) (*mesos.MesosSlave, error)
}
