package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"studygolang/chatroom/client/utils"
	"studygolang/chatroom/common/consts"
	common "studygolang/chatroom/common/message"
)

type FileProcess struct {

}

func (this *FileProcess) UploadBigFile(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	// 获取文件名
	split := strings.Split(path, "/")
	filename := split[len(split)-1:][0]

	uploadStart := true
	tf := &utils.Transfer{Conn: CurUser.Conn}
	reader := bufio.NewReader(file)

	// 循环 分段上传 (大文件，没法一次性全部上传，选择一次传一部分，直到全部文件内容传完)
	for {
		// 1024个byte = 1KB;   20 MB = 20*1024 KB = 20*1024*1024 byte
		var data [50 * 1024 * 1024]byte
		//var data [20]byte
		n, err := reader.Read(data[:])

		if n == 0 && err == io.EOF {
			_ = writeDataByNet(filename, data[0:0], 0, consts.UploadBigFileEnd, tf)
			fmt.Println("上传结束...")
			break
		}

		var err2 error
		if uploadStart {
			err2 = writeDataByNet(filename, data[:n], n, consts.UploadBigFileStart, tf)
			fmt.Println("上传开始...")
			uploadStart = false
		} else {
			err2 = writeDataByNet(filename, data[:n], n, consts.UploadBigFileing, tf)
			fmt.Println("上传中...")
		}

		if err2 != nil {
			fmt.Println("writeData error: ", err2)
		}
	}
	return
}

func writeDataByNet(filename string, content []byte, len int, status string, tf *utils.Transfer) (err error) {
	var mess common.Message
	mess.Type = common.FileUploadMesType

	var fileUpload common.FileUploadMes
	fileUpload.FileName = filename
	fileUpload.Data     = content
	fileUpload.Len      = len
	fileUpload.UploadStatus = status

	bytes, err := json.Marshal(fileUpload)
	if err != nil {
		return
	}

	mess.Data = string(bytes)
	data, err := json.Marshal(mess)
	if err != nil {
		return
	}

	return tf.WritePkg(data)
}
