package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	/*resp, err := http.NewRequest(
		http.MethodGet,
		url,
		nil)*/
		//resp.Header.

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("请求出错,status code: ", resp.StatusCode)
		return nil,
			fmt.Errorf("wrong status code; %d\n", resp.StatusCode)
	}

	// GBK to utf8
	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader,
		e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

	//all, err := ioutil.ReadAll(resp.Body)
	/*if err != nil {
		panic(err)
	}

	return nil, nil*/
}

func FetchUrl(url string) ([]byte, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")

	//resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	// 自定义 client
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("请求出错,status code: ", resp.StatusCode)
		return nil,
			fmt.Errorf("wrong status code; %d\n", resp.StatusCode)
	}

	// GBK to utf8
	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader,
		e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// 转换网页编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 缓存 io
	bytes, err := r.Peek(1024)
	if err != nil {
		//log.Error("Fetcher error: $v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
