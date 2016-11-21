package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func h(w http.ResponseWriter, req *http.Request) {
	buf, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal("request", err)
	}
	f, err := os.Create("/tmp/dat2")
	f.Write(buf)
}
