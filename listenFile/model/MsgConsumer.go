package model

import (
	"context"
	"encoding/json"
	"fmt"
)

type MsgConsumer struct {
	LinkTaskId int32
	MsgFrom chan *Msg
	Ctx context.Context
	CancelFunc context.CancelFunc
}

func (consumer *MsgConsumer) StartConsumer() error {
	go consumer.run()
	return nil
}

func (consumer *MsgConsumer) run() {
	exitTask := false
	for {
		select {
		case <-consumer.Ctx.Done():
			err := consumer.Cancel()
			if err != nil {
				fmt.Errorf("msgConsumer cancel failed! error: %v", err)
			}
			exitTask = true
		case msg, ok := <-consumer.MsgFrom:
			if !ok {
				fmt.Println("get msg from task failed!")
				break
			}

			marshal, _ := json.Marshal(msg)
			fmt.Printf("got msg from task: %s\n", string(marshal))
		}

		if exitTask {
			fmt.Printf("consumer[%d] process will exit\n", consumer.LinkTaskId)
			break
		}
	}
}

func(consumer *MsgConsumer) Cancel() error {
	// 进行任务取消前的收尾工作
	return nil
}