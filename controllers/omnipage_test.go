package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
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
