package parse

import (
	"engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	all := compile.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParseResult{}
	limit := 1
	for _, m := range all {
		// 城市
		result.Items = append(
			result.Items, "City_" + m[2])

		// URL
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: ParserCity,
			})
		limit --;
		if limit == 0 {
			break
		}
		//fmt.Printf("City: %s ; URL: %s\n", m[2], m[1]) // City: 上饶 ; URL: http://www.zhenai.com/zhenghun/shangrao
	}
	return result
}
