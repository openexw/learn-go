package parse

import (
	"engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var gender  = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var cityUrlRe = regexp.MustCompile(
	`href="(http://www.zhenai.com/zhenghun/[^"]+)`)

//const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParserCity(contents []byte) engine.ParseResult  {
	//compile := regexp.MustCompile(cityRe)
	all := cityRe.FindAllSubmatch(contents, -1)

	//fmt.Printf("----%v", all)
	result := engine.ParseResult{}
	// 获取性别
	gender := getGender(contents)
	//fmt.Printf("-=======%s", gender)
	for _, m := range all {
		// 城市
		name := string(m[2])
		result.Items = append(
			result.Items, "User_"+string(m[2]))

		// URL
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					//fmt.Printf("====%s-%s\n", name, gender)
					return ParserProfile(c, name, gender)
				},
			})
	}

	// city url
	matches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _,m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc:ParserCity,
			})
	}
	// todo 去重
	return result
}



func getDataByCityUrl (contents []byte) {

}
// 获取 性别
func getGender(contents []byte) string {
	//compile := regexp.MustCompile(gender)
	match := gender.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return  ""
	}
}