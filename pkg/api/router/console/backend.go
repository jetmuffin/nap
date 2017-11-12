package console

import (
	. "github.com/JetMuffin/nap/pkg/types"
)

// Backend abstracts a console manager
type Backend interface {
	GetTaskByID(taskId string) (MesosTask, error)
	TaskContainerName(task MesosTask) (string, error)
	TaskSlave(task MesosTask) (*MesosSlave, error)
}
