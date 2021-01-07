package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func BadRequst(res http.ResponseWriter) {
	var bytes bytes.Buffer;
	content := `
			Bad Request !
			<br />
			Request Schema is /:millisecond
			<br />
			eg. /6000
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

	regSecond := `/(\d+)`;
	matched, err := regexp.MatchString(regSecond, path);

	if err != nil || !matched  {
		BadRequst(res);
		return;
	}

	re, err := regexp.Compile(regSecond);
	if err != nil {
		fmt.Println(err);
		return;
	}
	match := re.FindAllStringSubmatch(path, -1);
	delay, _ := strconv.ParseInt(match[0][1], 10, 64);
	fmt.Println(path, delay, time.Duration(delay / 1000) * time.Second) 
	<- time.Tick(time.Duration(delay / 1000) * time.Second);
	res.Write([]byte(match[0][1]));	
}

func main() {
	http.HandleFunc("/", Handler);
	http.ListenAndServe(":6001", nil);
}
