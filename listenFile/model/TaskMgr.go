package model

import (
	"context"
	"studygolang/listenFile/config"
)

type TaskMgrModel struct {
	TaskMap map[int32]*Task
	MsgConsumerMap map[int32]*MsgConsumer
}

var TaskMgr *TaskMgrModel

func init()  {
	TaskMgr = &TaskMgrModel{
		TaskMap: map[int32]*Task{},
		MsgConsumerMap: map[int32]*MsgConsumer{},
	}
}

func (mgr *TaskMgrModel) CreateTask(filePath string) (*Task, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())

	task := &Task{
		TaskId:     config.GetTaskId(),
		FilePath:   filePath,
		Status:     config.TaskCreated,
		MsgChan:    make(chan *Msg, 64),
		Ctx:        ctx,
		CancelFunc: cancelFunc,
	}
	mgr.TaskMap[task.TaskId] = task

	ctx2, cancelFunc2 := context.WithCancel(context.Background())
	consumer := &MsgConsumer{
		LinkTaskId: task.TaskId,
		MsgFrom:    task.MsgChan,
		Ctx:        ctx2,
		CancelFunc: cancelFunc2,
	}
	mgr.MsgConsumerMap[task.TaskId] = consumer

	return task, nil
}

func (mgr *TaskMgrModel) RunTask(taskId int32) (err error) {
	err = mgr.TaskMap[taskId].ListenFile()
	if err != nil {
		return
	}

	err = mgr.MsgConsumerMap[taskId].StartConsumer()
	if err != nil {
		return
	}
	return nil
}

func (mgr *TaskMgrModel) CancelTask(task *Task) error {
	// 取消任务
	task.CancelFunc()
	mgr.MsgConsumerMap[task.TaskId].CancelFunc()

	delete(mgr.TaskMap, task.TaskId)
	delete(mgr.MsgConsumerMap, task.TaskId)
	return nil
}

func (mgr *TaskMgrModel) ExitAllTaskProcess() error {
	for _, task := range mgr.TaskMap {
		task.CancelFunc()
	}
	for _, msgConsumer := range mgr.MsgConsumerMap {
		msgConsumer.CancelFunc()
	}
	return nil
}

