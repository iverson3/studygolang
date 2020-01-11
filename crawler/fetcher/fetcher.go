package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		//return nil, errors.New("Error: status code")
	}

	// 重新创建一个utf8编码格式的Reader 将其他编码格式的网页转为utf8
	// 但下面代码目前无法使用  因为对应的包还没下载

	// 如果确定知道网页的编码方式为 gbk  则可以直接进行转换  否则就需要先获取判断网页的编码方式
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	// 先获取到当前网页的编码方式 再进行编码转换为utf8
	//bodyReader := bufio.NewReader(resp.Body)
	//e := determineEncoding(bodyReader)
	//utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(resp.Body)
}

//func determineEncoding(r *bufio.Reader) encoding.Encoding {
//	// 从中取1024字节的内容出来
//	bytes, err := r.Peek(1024)
//	if err != nil {
//		return unicode.UTF8
//	}
//	// 判断确定内容的编码方式
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}





































