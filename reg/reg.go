package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is 1065173619@qq.com
ww www@qq.com
sdd@qq.com`

func main() {
	//compile, e := regexp.Compile("1065173619@qq.com")
	compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\\.([a-zA-Z0-9]+)`)
	//match := compile.FindAllString(text, -1)
	match := compile.FindAllStringSubmatch(text, -1)
	for _,v := range match{
		fmt.Printf("%s@%s.%s", v[0], v[1], v[2])
	}
}