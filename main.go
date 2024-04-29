package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	http.HandleFunc("/post", testPost)
	http.HandleFunc("/conf", testConf)
	http.HandleFunc("/hello", testHello)

	http.ListenAndServe(":8889", nil)
}

func testHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!\n"))
	println("hello world")
}

// test post request
func testPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := os.Mkdir("data", 0755)
		if os.IsExist(err) {
			println("data path is exist.")
		}
		fd, err := os.OpenFile("data/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			println("Open data/test.txt file error: ", err.Error())
			return
		}
		defer fd.Close()
		fd.Write([]byte("this is a test\n"))
		println("test post")
	}
}

// test config
func testConf(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ymlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			println("Read config file error")
			return
		}
		var config Config
		err = yaml.Unmarshal(ymlFile, &config)
		if err != nil {
			println("Unmarshal config file error")
			return
		}
		println("name: ", config.Name)
		println("addr: ", config.Addr)
	}
}

type Config struct {
	Name string
	Addr string
}
