/*package main

import "interface/mock/MockRetiever"
type Retriever interface {
	Get(url string) string
}
func download(r Retriever)  string {
	return r.Get("http://www.baidu.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{"this is a baidu.com"}

	download(r)
}*/

package main

import (
	"fmt"
)

/**
手机接口
 */
type Phone interface {
	call()
}

/**
华为手机
 */
type HuaweiPhone struct {
	Contents string
}
/**
华为手机实现 Phone 的call()
 */
func (hw HuaweiPhone) call() {
	fmt.Println(hw.Contents)
}


/**
小米手机 Phone 的call()
 */
type MiPhone struct {
	Contents string
}
/**
小米手机的实现
 */
func (mi MiPhone) call() {
	fmt.Println(mi.Contents)
}


func main() {
	//io.WriteCloser()
	hw := HuaweiPhone{"this is Huawei Phone."}
	hw.call()

	//fmt.Printf("%T %v\n", hw)
	mi := MiPhone{"this is xiaomi Phone."}
	mi.call()
}

