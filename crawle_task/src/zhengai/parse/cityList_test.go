package parse

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 避免网址没法打开的情况，先将文件保存在本地，然后从本地获取数据
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("city_list_test_data.html")
	if err != nil {
		panic(err)
	}

	cityList := ParseCityList(contents)
	fmt.Printf("%v", cityList)
	//fmt.Printf("%s", contents)
	// 验证有 470 个城市
	const resultSize = 470
	urls  := []string {
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/baicheng1",
		"http://www.zhenai.com/zhenghun/cangzhou",
	}
	cities := []string {
		"City_阿坝",
		"City_白城",
		"City_沧州",
	}
	if  len(cityList.Items) != resultSize {
		t.Errorf("应该有 %d 个城市，但是只有 %d 个城市",
			resultSize, len(cityList.Items))
	}

	for i, city := range cities {
		if cityList.Items[i].(string) != city {
			t.Errorf("url #%d, %s, but was %s",
				i, city, cityList.Items[i].(string))
		}
	}
	// 验证 470 个URL
	if resultSize != len(cityList.Requests){
		t.Errorf("应该有 %d 个Url，但是只有 %d 个Url",
			resultSize, len(cityList.Requests))
	}
	for i, url := range urls {
		if cityList.Requests[i].Url != url {
			t.Errorf("url #%d, %s, but was %s",
				i, url, cityList.Requests[i].Url)
		}
	}
}