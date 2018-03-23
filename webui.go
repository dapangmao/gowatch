package main

import (
	"net/http"
	"golang.org/x/tools/go/gcimporter15/testdata"
	"reflect"
	"strconv"
)



func console(w http.ResponseWriter, r *http.Request) {

}

func gowatch(port ...interface{}) {
	var res string
	var p interface{}
	if len(port) == 1 {
		p = port[0]
	}

	if p == nil {
		res = "8011"
	} else if _p, ok := p.(int); ok {
		res = strconv.Itoa(_p)
	} else if _p, ok := p.(string); ok {
		res = _p
	}

	http.HandleFunc("/", console)
	http.ListenAndServe(":"+res, nil)

}
