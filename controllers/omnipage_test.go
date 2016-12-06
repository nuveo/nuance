package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/nuveo/nuance/config"
	"github.com/nuveo/nuance/omnipage"
)

func startOmnipage() {
	cfg := config.Nuance{}
	env.Parse(&cfg)

	op := omnipage.New()

	err := op.SetLicense(cfg.OemLicenseFile, cfg.OemCode)
	if err != nil {
		log.Fatal("SetLicense failed:", err)
	}

	err = op.Init(cfg.CompanyName, cfg.ProductName)
	if err != nil {
		log.Fatal("Init failed:", err)
	}

	SetOmnipage(&op)
	SetConfig(&cfg)
}

func TestOcrFile(t *testing.T) {
	startOmnipage()

	txt, err := ocrFile("../test_assets/testFile.png")
	if err != nil {
		t.Fatal("ocrFile:", err)
	}
	txt = strings.Replace(txt, "  ", " ", -1)
	if !strings.Contains(txt, "It is a test.") {
		t.Error("Error: string does not contain expected substring.")
	}
}

func TestTotextApplicationJson(t *testing.T) {
	startOmnipage()

	router := mux.NewRouter()
	router.HandleFunc("/omnipage/totext", ImgToText).Methods("POST")
	server := httptest.NewServer(router)
	defer server.Close()

	pathPth := "../test_assets/testFile.png"

	buff, err := ioutil.ReadFile(pathPth)
	if err != nil {
		t.Fatal("ReadFile:", err)
		return
	}

	r := request{}
	r.Base64 = base64.StdEncoding.EncodeToString(buff)

	jsonBuffer := new(bytes.Buffer)
	json.NewEncoder(jsonBuffer).Encode(r)

	res, err := http.Post(server.URL+"/omnipage/totext",
		"application/json; charset=utf-8",
		jsonBuffer)
	if err != nil {
		t.Fatal("POST:", err)
		return
	}

	//io.Copy(os.Stdout, res.Body)
	decoder := json.NewDecoder(res.Body)

	var rp response
	err = decoder.Decode(&rp)
	if err != nil {
		t.Fatal("Decode:", err)
	}

	//fmt.Println(rp.Text)
	txt := strings.Replace(rp.Text, "  ", " ", -1)
	if !strings.Contains(txt, "It is a test.") {
		t.Error("Error: string does not contain expected substring.")
	}

}

func TestTotextMultipartFormData(t *testing.T) {
	startOmnipage()

	router := mux.NewRouter()
	router.HandleFunc("/omnipage/totext", ImgToText).Methods("POST")
	server := httptest.NewServer(router)
	defer server.Close()

	files := []string{"../test_assets/testFile.png"}

	reader, err := post(server.URL+"/omnipage/totext", files)
	if err != nil {
		t.Fatal("reader:", err)
	}

	//io.Copy(os.Stdout, reader)
	decoder := json.NewDecoder(reader)

	var rp response
	err = decoder.Decode(&rp)
	if err != nil {
		t.Fatal("Decode:", err)
	}

	//fmt.Println(rp.Text)
	txt := strings.Replace(rp.Text, "  ", " ", -1)
	if !strings.Contains(txt, "It is a test.") {
		t.Error("Error: string does not contain expected substring.")
	}
}

func post(url string, files []string) (bodyReader io.Reader, err error) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for k, v := range files {
		var f *os.File
		f, err = os.Open(v)
		if err != nil {
			return
		}

		var fw io.Writer
		fw, err = w.CreateFormFile("image_"+string(k), v)
		if err != nil {
			return
		}

		if _, err = io.Copy(fw, f); err != nil {
			return
		}
		f.Close()
	}

	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}

	//io.Copy(os.Stdout, res.Body)

	//var bAux []byte
	//bAux, err = ioutil.ReadAll(res.Body)

	//txt = string(bAux)
	//res.Body.Close()

	bodyReader = res.Body
	return
}
