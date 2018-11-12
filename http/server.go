package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func mobileRequest(url string)  {
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	//resp, err := http.DefaultClient.Do(request)

	// 自定义 client
	client := http.Client{
		CheckRedirect: func(req *http.Request,
			via []*http.Request) error {
			fmt.Println("------------------r------------", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}

// 基本请求
func baseRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}

func main()  {

	mobileRequest("https://www.baidu.com")
}
