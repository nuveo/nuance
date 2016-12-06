package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nuveo/nuance/config"
	"github.com/nuveo/nuance/omnipage"
)

var op *omnipage.Omnipage
var cfg *config.Nuance

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type request struct {
	Base64 string
}

type response struct {
	Text string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SetConfig(c *config.Nuance) {
	cfg = c
}

func SetOmnipage(opInstance *omnipage.Omnipage) {
	op = opInstance
}

func ImgToText(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]

	if contentType == "application/json" {
		decoder := json.NewDecoder(r.Body)

		var jr request
		err := decoder.Decode(&jr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buff := []byte{}
		buff, err = base64.StdEncoding.DecodeString(jr.Base64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := cfg.TmpPath + "/omnipage_" + randString(20)

		err = ioutil.WriteFile(filename, buff, 0644)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		txt := ""
		txt, err = ocrFile(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.Remove(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := response{}
		resp.Text = txt

		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		fmt.Fprint(w, string(b))

	} else if contentType == "multipart/form-data" {

		reader, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		txt := ""
		for {
			var part *multipart.Part
			part, err = reader.NextPart()
			if err == io.EOF {
				break
			}

			if part.FileName() == "" {
				continue
			}

			log.Println("filename", part.FileName())

			filename := cfg.TmpPath + "/omnipage_" + randString(20)

			var dst *os.File
			dst, err = os.Create(filename)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if _, err = io.Copy(dst, part); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			txtAux := ""
			txtAux, err = ocrFile(filename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = os.Remove(filename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			txt += txtAux
		}

		resp := response{}
		resp.Text = txt

		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		fmt.Fprint(w, string(b))

	} else {

		errMsg := "Content-Type: \"" + contentType + "\" not supported"
		log.Println("Content-Type", contentType)
		http.Error(w, errMsg, http.StatusBadRequest)
	}
}

func ocrFile(fullPath string) (txt string, err error) {

	op.SetLanguagePtBr() // TODO: implement SetLanguage REST interface
	op.SetCodePage("UTF-8")

	txt, err = op.OCRImgToText(fullPath)
	if err != nil {
		return
	}

	return
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
