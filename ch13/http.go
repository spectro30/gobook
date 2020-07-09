package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "text/html")
	//http://127.0.0.1:9000/hello
	io.WriteString(
		res,
		"Hello World",
	)

}

func main()  {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}