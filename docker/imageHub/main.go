package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 镜像仓库服务

const (
	DefaultStorePath = "/usr/local/xdocker/images/"
)

type ImageInfo struct {
	Name string
	Tag string
	Size int64
	CreateTime int64
}

func init() {
	exist, err := PathExist(DefaultStorePath)
	if err != nil {
		panic(err)
	}
	if !exist {
		// 不存在则创建目录
		err = os.MkdirAll(DefaultStorePath, 0666)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	http.HandleFunc("/images/push", pushImage)
	http.HandleFunc("/images/list", listImage)
	http.HandleFunc("/images/search", searchImage)
	http.HandleFunc("/images/pull", pullImage)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}



func searchImage(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")

	imageInfoList, err := GetAllImageDir()
	if err != nil {
		w.Write([]byte("get image list fail"))
		return
	}

	images := make([]string, 0)
	for _, image := range imageInfoList {
		// 判断镜像名是否包含搜索的关键字
		if strings.Contains(image.Name, keyword) {
			images = append(images, image.Name)
		}
	}

	bytes, err := json.Marshal(images)
	if err != nil {
		w.Write([]byte("fail"))
		return
	}

	w.Write(bytes)
}

func listImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("imagename")

	var err error
	var imageInfoList []*ImageInfo
	if imageName == "" {
		imageInfoList, err = GetAllImageDir()
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("get image list fail"))
			return
		}
	} else {
		// 获取指定镜像名目录下的所有镜像文件 (相同的镜像 不同的tag)
		imageInfoList, err = GetImageListWithName(imageName)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("get image list fail"))
			return
		}
	}

	images := make([]string, 0, len(imageInfoList))
	for _, image := range imageInfoList {
		images = append(images, image.Name)
	}

	bytes, err := json.Marshal(images)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("fail"))
		return
	}

	w.Write(bytes)
}

func pushImage(w http.ResponseWriter, r *http.Request) {
	// 获取镜像名和tag
	imageName := r.FormValue("imagename")
	tag := r.FormValue("tag")

	// 判断对应的镜像是否已经存在
	exist, _, err := ImageIsExist(imageName, tag)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("find imageFile fail"))
		return
	}
	if exist {
		w.Write([]byte("image exist"))
		return
	}

	// 从请求中获取上传的镜像文件
	formFile, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("get file fail"))
		return
	}
	defer formFile.Close()

	// 镜像目录： /usr/local/xdocker/images/xxx/
	imageDirPath := fmt.Sprintf("%s%s/", DefaultStorePath, imageName)
	// 判断该镜像目录是否存在，不存在则创建目录
	pathExist, err := PathExist(imageDirPath)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("find imageDir fail"))
		return
	}
	if !pathExist {
		err = os.MkdirAll(imageDirPath, 0666)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("create imageDir fail"))
			return
		}
	}

	// 镜像文件保存路径： /usr/local/xdocker/images/xxx/xxx@1.2.0.tar
	dstFile, err := os.OpenFile(fmt.Sprintf("%s%s@%s.tar", imageDirPath, imageName, tag), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("create dst file fail"))
		return
	}
	defer dstFile.Close()

	// 将post提交上来的文件拷贝保存到目标文件中
	_, err = io.Copy(dstFile, formFile)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("copy file fail"))
		return
	}

	w.Write([]byte("ok"))
}

func pullImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("decode fail"))
		return
	}

	// 获取镜像名和tag
	imageName := params["imagename"]
	tag := params["tag"]

	fmt.Println(imageName, tag)

	// 检查是否存在该镜像
	exist, imageFilePath, err := ImageIsExist(imageName, tag)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("find imageFile fail"))
		return
	}
	if !exist {
		w.Write([]byte("image not exist"))
		return
	}

	imageFile, err := os.Open(imageFilePath)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("open target file fail"))
		return
	}
	defer imageFile.Close()

	stat, err := imageFile.Stat()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("get fileinfo fail"))
		return
	}

	w.Header().Set("Content-Type", "application/x-zip-compressed")
	//也可用http.DetectContentType获得文件类型：application/zip
	w.Header().Set("Content-Disposition", "attachment; filename=" + stat.Name())
	w.Header().Set("Content-Length", strconv.FormatInt(stat.Size(), 10))

	n, err := io.Copy(w, imageFile)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("copy file fail"))
		return
	}
	if n != stat.Size() {
		w.Write([]byte("copy file length is wrong"))
		return
	}
}

// ImageIsExist 判断镜像是否存在，存在则返回该镜像文件的完整路径
func ImageIsExist(imageName, tag string) (bool, string, error) {
	imageDirPath := fmt.Sprintf("%s%s/", DefaultStorePath, imageName)
	fmt.Println(imageDirPath)
	// 先判断该镜像的存储目录是否存在
	exist, err := PathExist(imageDirPath)
	if err != nil {
		return false, "", err
	}
	if !exist {
		return false, "", nil
	}

	// 遍历镜像目录
	imageFileList, err := ioutil.ReadDir(imageDirPath)
	if err != nil {
		return false, "", err
	}

	var imageExist bool
	var imageFullPath string
	targetFileName := fmt.Sprintf("%s@%s.tar", imageName, tag)
	for _, imageFile := range imageFileList {
		if imageFile.Name() == targetFileName {
			imageExist = true
			imageFullPath = fmt.Sprintf("%s%s", imageDirPath, targetFileName)
			break
		}
	}
	fmt.Println(imageFullPath)
	return imageExist, imageFullPath, nil
}

func GetImageListWithName(imageName string) ([]*ImageInfo, error) {
	imageDirPath := fmt.Sprintf("%s%s/", DefaultStorePath, imageName)

	list := make([]*ImageInfo, 0)
	// 判断该镜像目录是否存在
	exist, err := PathExist(imageDirPath)
	if err != nil {
		return nil, err
	}
	if !exist {
		return list, nil
	}

	// 遍历镜像目录下所有的镜像文件
	dirList, err := ioutil.ReadDir(imageDirPath)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirList {
		nameArr := strings.Split(dir.Name(), ".")
		name := strings.Join(nameArr[:len(nameArr) - 1], ".")

		list = append(list, &ImageInfo{Name: name})
	}
	return list, nil
}

func GetAllImageDir() ([]*ImageInfo, error) {
	dirList, err := ioutil.ReadDir(DefaultStorePath)
	if err != nil {
		return nil, err
	}
	if len(dirList) == 0 {
		return nil, nil
	}

	imageList := make([]*ImageInfo, 0)
	// 遍历目录下的文件夹 (每个文件夹都是一个镜像目录)
	for _, dir := range dirList {
		// 文件夹名即是镜像名，但具体的镜像文件在文件夹下
		// 不同tag的相同镜像，都会放在同一个目录下
		imageList = append(imageList, &ImageInfo{
			Name: dir.Name(),
		})
	}

	return imageList, nil
}

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}