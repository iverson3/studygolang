package process

import (
	"fmt"
	"studygolang/listenFile/config"
	"studygolang/listenFile/model"
)

func Start() {
	for {
		var key int
		var filePath string
		var taskId int
		quit := false

		fmt.Println("==========可选项 =========")
		fmt.Println("        1 - 新增一个任务")
		fmt.Println("        2 - 删除一个任务")
		fmt.Println("        3 - 列出当前任务")
		fmt.Println("        4 - 退出")
		fmt.Println("=========================")

		_, _ = fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("请输入文件路径：")
			_, _ = fmt.Scanf("%s\n", &filePath)

			task, err := model.TaskMgr.CreateTask(filePath)
			if err != nil {
				fmt.Errorf("create task failed，error: %v", err)
				break
			}
			fmt.Println("create task success!")

			err = model.TaskMgr.RunTask(task.TaskId)
			if err != nil {
				fmt.Errorf("task run failed! error: %v", err)
				break
			}
			fmt.Println("task run success!")
		case 2:
			for {
				fmt.Println("请输入删除任务的Id：")
				_, _ = fmt.Scanf("%d\n", &taskId)

				task, ok := model.TaskMgr.TaskMap[int32(taskId)]
				if !ok {
					fmt.Errorf("输入的任务Id有误!：")
					continue
				}

				err := model.TaskMgr.CancelTask(task)
				if err != nil {
					fmt.Errorf("cancel task failed! error: %v", err)
					break
				}

				fmt.Println("cancel Task success!")
				break
			}
		case 3:
			fmt.Println("task List:")
			if len(model.TaskMgr.TaskMap) == 0 {
				fmt.Println("no task")
				break
			}
			for _, task := range model.TaskMgr.TaskMap {
				fmt.Printf("%d  %s  %s \n", task.TaskId, task.FilePath, config.TaskStatusMap[task.Status])
			}
			fmt.Println()
		case 4:
			quit = true
		default:
			fmt.Println("选择有误，请重新进行选择")
		}

		if quit {
			fmt.Println("bye~")
			break
		}
	}
}
