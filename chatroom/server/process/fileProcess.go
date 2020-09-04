package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"studygolang/chatroom/common/consts"
	common "studygolang/chatroom/common/message"
)

var (
	UploadDir = "d:/"
)

type FileProcess struct {

}

func (this *FileProcess) UploadBigFile(mess *common.Message) (err error) {
	var uploadFile common.FileUploadMes
	err = json.Unmarshal([]byte(mess.Data), &uploadFile)
	if err != nil {
		return
	}

	// 拼装文件保存路径
	filePath := UploadDir + "target_" + uploadFile.FileName

	// 大文件上传开始
	if uploadFile.UploadStatus == consts.UploadBigFileStart {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		return saveDataToFile(file, uploadFile)
	} else {
		// 上传结束
		if uploadFile.UploadStatus == consts.UploadBigFileEnd {
			// 关闭文件
			fmt.Println("上传完成")
			return nil
		}

		// 大文件上传中
		file, err := os.OpenFile(filePath, os.O_APPEND, 666)
		if err != nil {
			return err
		}
		return saveDataToFile(file, uploadFile)
	}
}

func saveDataToFile(file *os.File, uploadFile common.FileUploadMes) (err error) {
	defer file.Close()

	writer := bufio.NewWriter(file)
	n, err := writer.Write(uploadFile.Data)
	if n != uploadFile.Len || err != nil {
		return
	}
	return writer.Flush()
}