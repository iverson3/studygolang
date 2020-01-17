package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"studygolang/crawler_distributed/config"
	"time"
)

// 使用定时器channel限制请求的频率
// 减小给别人服务器的压力，防止被封客户端ip
var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	log.Printf("Fetching url %s", url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 设置请求header信息
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	request.Header.Add("cookie", "sid=65044270-f381-4ba3-a576-e82993c94fa4; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1578900349; FSSBBIl1UgzbN7N443S=17Zvbuq_CR4CDes7GzxPyjDb2Zyk0irlOt72XW64n3BX9RJd0sI20_mXqDoRQ1Ki; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1578904462; FSSBBIl1UgzbN7N443T=4NJoQpQu0Hthw5pDUnWG45hTX_knxfQmqLo6E7EGaoAZdZAEDi_0xE7EG69qSME7Y6ahBGA99cd9UZri0j6t31lUWVBvRmCmh7fQSjzAk1ZE78hNv2JNV1rEgPdw4ZiRJysYAde6Buv7Go_74gGTRewb1B4kFbFuMEsTnsDHHElPsJc8scYb0J57FTiI.my5JBPoS2indJEIN.AYaiLgGM8iw2LCf4H9IXW30zcUfcsoEYKulsCs5HJ5o0yMzmLpprGtWcRpETEWp22.U5vpHVVFXtUYTMKMcCfvupaiGIsk1KZ0_d5wKFSTRMg7qNec5vtPrjIMCOlxrok93fkARYMV5joQaAMpVsLDsmevKGWJiAW7krQRiAjk_wWAGypG8_5g")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		//return nil, errors.New("Error: status code")
	}
	return ioutil.ReadAll(resp.Body)
}

func Fetch2(url string) ([]byte, error) {
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





































