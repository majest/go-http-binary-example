package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, req *http.Request) {
	buf, err := ioutil.ReadAll(req.Body)
	fname := req.URL.Query().Get("name")

	if err != nil || fname == "" {
		w.Write([]byte("error: No file"))
		return
	}

	f, err := os.Create(fmt.Sprintf("/tmp/data/%s", fname))

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	f.Write(buf)
	w.Write([]byte("ok"))
}

func main() {
	fmt.Println("Starting http file sever")
	http.HandleFunc("/upload", upload)
	http.Handle("/", http.FileServer(http.Dir("/tmp/data")))

	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println(err)
	}
}
