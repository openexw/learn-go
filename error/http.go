package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net/http"
	"os"
)

func server(w http.ResponseWriter,
	r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
		//panic(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	w.Write(all)
	return nil
}

type appHandler func(w http.ResponseWriter, r *http.Request) error
/**
统一错误处理
 */
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request)  {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request； %s", err.Error())
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err): // 文件不存在
				code = http.StatusNotFound
				/*http.Error(writer,
					http.StatusText(http.StatusNotFound),
					http.StatusNotFound)*/
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
					http.StatusText(code),
					code)
		}
	}
}
func hello (w http.ResponseWriter,
	r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)

	//log.Warn("%s", path)
	if err != nil {
		return err
		//panic(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	w.Write(all)
	return nil
}
func main() {
	//http.HandlerFunc("/list/", server)

	http.HandleFunc("/list", errWrapper(hello))
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
