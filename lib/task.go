package lib

import (
	"scu/main/lib/cache"
)

type Task struct {
	ws_name   string
	task_name string
	status    int
	Deps      []string
	Run       task_run
}

type task_run = func(c *cache.Cache)

const (
	TASK_STATUS_PENDING = iota
	TASK_STATUS_RUNNING = iota
	TASK_STATUS_SUCCESS = iota
	TASK_STATUS_FAILURE = iota
)

func NewTask(ws_name string, task_name string, deps []string, run task_run) Task {
	return Task{
		ws_name:   ws_name,
		task_name: task_name,
		status:    TASK_STATUS_PENDING,
		Deps:      deps,
		Run:       run,
	}
}
