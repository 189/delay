package main

import (
	"bytes"
	"delay/Help"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func BadRequst(res http.ResponseWriter) {
	var bytes bytes.Buffer;
	content := `
			Bad Request !
			<br />
			Request Schema is /:second
			<br />
			eg. /6
			<br />
			will got response after 6 second
		`;
		bytes.WriteString(content);
		res.Write(bytes.Bytes());
}

func Handler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path;

	res.Header().Set("Content-Type", "text/html; charset=utf-8");
	
	if path == "/" {
		BadRequst(res);
		return;
	}

	delay, err := Help.GetSecond(path);
	
	if err != nil {
		fmt.Println(err);
		res.Write([]byte(err.Error()));
		return;
	}

	fmt.Println(path, delay, time.Duration(delay) * time.Second);
	<- time.Tick(time.Duration(delay) * time.Second);

	res.Write([]byte(strconv.FormatInt(delay, 10)));
}

func main() {
	http.HandleFunc("/", Handler);
	http.ListenAndServe(":6001", nil);
	// start := time.Now();
	// fmt.Println(start);
	// t := <-time.Tick(0 * time.Second);
	// fmt.Println(t, "\n", start.Sub(time.Now()))
}
