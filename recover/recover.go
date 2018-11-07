package main

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			println(err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("JJ"))

	a := 5 /0
	println(a)
}
func main() {
	tryRecover()
}
