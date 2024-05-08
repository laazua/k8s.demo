// 
// 仅用于测试k8s资源对象的代码
//

package main

import (
        "encoding/json"
        "fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Book struct {
	Name   string
	Author string
}

func main() {
	router := mux.NewRouter()
        router.HandleFunc("/test", healthTest).Methods("GET")
	router.HandleFunc("/hello/{name}", helloDemo).Methods("GET")
	router.HandleFunc("/tasks/{id}", tasksDemo).Methods("GET", "POST")
	router.HandleFunc("/books/info", booksInfo).Methods("GET")
	router.HandleFunc("/books/{item}", booksItem).Methods("GET")
        router.HandleFunc("/books/check", booksCheck).Methods("POST")

        println("Listen On: 8889")
	http.ListenAndServe(":8889", router)
}

func healthTest(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Health Test Success ..."))
}

func helloDemo(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte("hello " + name))
}

func tasksDemo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Write([]byte("task " + id))
}

// test current path file
func booksItem(w http.ResponseWriter, r *http.Request) {
	item := mux.Vars(r)["item"]
	ymlFile, err := os.ReadFile("book.yaml")
	if err != nil {
		w.Write([]byte("Read file error"))
		return
	}
	var book Book
	err = yaml.Unmarshal(ymlFile, &book)
	if err != nil {
		w.Write([]byte("Unmarshal book.yaml error"))
		return
	}
	if item == "name" {
		w.Write([]byte("Name is " + book.Name))
		return
	}
	if item == "author" {
		w.Write([]byte("Author is " + book.Author))
		return
	}
	w.Write([]byte("Wrong Item of book"))
}

// test config path file
func booksInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ymlFile, err := os.ReadFile("config/book.yaml")
	if err != nil {
		w.Write([]byte("Read file error"))
		return
	}
	var book Book
	err = yaml.Unmarshal(ymlFile, &book)
	if err != nil {
		w.Write([]byte("Unmarshal config file error"))
		return
	}
	w.Write([]byte("Book's introduction: "))
        // 返回json数据
	json.NewEncoder(w).Encode(book)
}

func booksCheck(w http.ResponseWriter, r *http.Request) {
        var book Book
        // 获取请求体的json数据
        json.NewDecoder(r.Body).Decode(&book)
        fmt.Fprintf(w, "name: %s, author: %s\n", book.Name, book.Author)
}
