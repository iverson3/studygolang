package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 设置请求header信息
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.16 Safari/537.36")
	request.Header.Add("cookie", "sid=c9bdd77e-69e5-4618-b16f-cbe38c2ed6a4; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1578736872; FSSBBIl1UgzbN7N443S=BX.KtTrlaiNToScm_Shyn74kyknlfWuMGSyYemn8mYPXIuuCsmjk5TP_tTPljGse; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1578843370; FSSBBIl1UgzbN7N443T=4RqaEZ3zXJ.s7oAAXcngmCZ_397iC6GznNgwF76dh5aOAY.qhEVlMptuCl2ZrnaTf5y3CqBUD0Wf_DUFDoYiwuVhjbBnnASBQwgMAIOcfBnugNJkVPm1X2urphFg52p0Iw577EEjXuS.szvM.sjnZWahdLG7DokrEncnJUVOTxV6pldm6oOn1j3.vCP_kEbe1b1_DuqjYAsyIrb4gGPqL3a_9JSafzmOdDXVPJVXzEkZyYLSJdqYZG_v9cEuPJWGrVpI3zXit2te9nIziyYZISApU0UAcOzmb_8F4DDEv.2ntYVWWxzAurALX0bkct439180LhaBc8Bk.I3nB03mgzpNH")

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





































