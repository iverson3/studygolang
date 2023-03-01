package util

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var defaultHttpClient *http.Client

type httpClient struct {
	conn *http.Client
}

func init() {
	//tr := &http.Transport{
	//	//MaxIdleConns:           0,
	//	MaxIdleConnsPerHost:    500,
	//	//MaxConnsPerHost:        2,
	//	//IdleConnTimeout:        0,
	//	//ForceAttemptHTTP2:      false,
	//	//DialContext: (&net.Dialer{
	//	//	Timeout: 2 * time.Second,   // tcp链接建立的超时时间
	//	//}).DialContext,
	//}
	//c := &http.Client{Transport: tr}
	//
	//defaultHttpClient = c
	defaultHttpClient = http.DefaultClient
	defaultHttpClient.Timeout = 10 * time.Second
}

func NewHttpClient(conn *http.Client, timeout int) *httpClient {
	if conn == nil {
		conn = defaultHttpClient
	}
	if timeout != 0 {
		conn.Timeout = time.Duration(timeout) * time.Second
	}
	return &httpClient{conn: conn}
}

func (hc *httpClient) GET(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}
	resp, err := hc.conn.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.StatusCode: %s", resp.Status)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	return bytes, err
}
