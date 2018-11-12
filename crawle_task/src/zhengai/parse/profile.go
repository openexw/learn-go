package parse

import (
	"engine"
	"model"
	"regexp"
	"strconv"
)

var ageRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var marriageRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
var heightRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
var weightRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
var incomeRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
var educationRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
var occupationRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>工作地:([^<]+)</div>`)
//var houseRe  = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
//var educationRe  = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)
//var nameRe = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)

func ParserProfile(contents []byte,
	name string,
	gender string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Gender = gender

	/*match  := ageRe.FindSubmatch(contents)
	if match != nil {

		if err != nil {
			//user age is age
			profile.Age = age
		}
	}*/
	/*age, err := strconv.Atoi(extractString(contents,
		ageRe))*/
	//age := extractString(contents, ageRe)
	age, _ := strconv.Atoi(extractString(contents, ageRe))

	profile.Age = age
	// 身高
	height, _ := strconv.Atoi(extractString(contents,
		heightRe))
	profile.Height = height
	// 体重
	weight, _ := strconv.Atoi(extractString(contents,
		weightRe))
	profile.Weight = weight
	// 婚姻状况
	//reg = regexp.MustCompile(marriageRe)
	/*match  = marriageRe.FindSubmatch(contents)
	if match != nil {
		profile.Marriage = string(match[1])
	}*/
	profile.Marriage = string(extractString(contents, marriageRe))
	profile.Income = string(extractString(contents, incomeRe))
	profile.Education = string(extractString(contents, educationRe)) // -- 有问题
	profile.Occupation = string(extractString(contents, occupationRe))

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return  result
}


func extractString(contents []byte, re *regexp.Regexp) string  {
	match  := re.FindSubmatch(contents)

	if len(match) == 2 {
		return string(match[1])
	} else if len(match) >= 2 {
		return string(match[2])
	} else {
		return  ""
	}

}