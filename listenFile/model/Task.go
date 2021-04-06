package model

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"studygolang/listenFile/config"
	"time"
)

type Task struct {
	TaskId int32
	FilePath string
	Status int
	Tail *tail.Tail
	MsgChan chan *Msg
	Ctx context.Context
	CancelFunc context.CancelFunc
}

func (task *Task) ListenFile() error {
	// 任务开始运行之前的逻辑
	conf := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Follow:      true,
	}

	tails, err := tail.TailFile(task.FilePath, conf)
	if err != nil {
		return err
	}
	task.Tail = tails

	go task.run()
	return nil
}

func(task *Task) run() {
	task.Status = config.TaskRunning
	exitTask := false
	for {
		select {
		case <-task.Ctx.Done():
			task.Status = config.TaskCanceled
			err := task.Cancel()
			if err != nil {
				fmt.Errorf("task cancel failed! error: %v\n", err)
			}
			exitTask = true
		case msg, ok := <-task.Tail.Lines:
			if !ok {
				fmt.Printf("no message from file: %s\n", task.FilePath)
				time.Sleep(1 * time.Second)
				break
			}
			// 处理任务
			m := &Msg{
				Content:    msg.Text,
				Len:        len(msg.Text),
				ModifyTime: msg.Time,
			}

			fmt.Println("send msg to consumer")
			task.MsgChan<- m
		}

		if exitTask {
			fmt.Printf("task[%d] process will exit\n", task.TaskId)
			break
		}
	}
}

func(task *Task) Cancel() error {
	// 进行任务取消前的收尾工作
	return nil
}