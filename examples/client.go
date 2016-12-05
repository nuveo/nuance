package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type request struct {
	Base64 string
}

func main() {

	pathPth := flag.String("file", "", "file to OCR")
	flag.Parse()

	buff, err := ioutil.ReadFile(*pathPth)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := request{}
	r.Base64 = base64.StdEncoding.EncodeToString(buff)

	jsonBuffer := new(bytes.Buffer)
	json.NewEncoder(jsonBuffer).Encode(r)

	res, err := http.Post("http://localhost:4000/omnipage/totext",
		"application/json; charset=utf-8",
		jsonBuffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, res.Body)
}
