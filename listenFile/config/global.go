package config

import "sync/atomic"

var (
	TaskId int32 = 0
)

const (
	TaskCreated = 0
	TaskRunning = 1
	TaskPaused = 2
	TaskCanceled = 3
)

var TaskStatusMap = make(map[int]string)

func init()  {
	TaskStatusMap[TaskCreated]  = "创建"
	TaskStatusMap[TaskRunning]  = "运行中"
	TaskStatusMap[TaskPaused]   = "暂停"
	TaskStatusMap[TaskCanceled] = "取消"
}

func GetTaskId() int32 {
	return atomic.AddInt32(&TaskId, 1)
}