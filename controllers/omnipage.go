package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/nuveo/nuance/omnipage"
)

var op *omnipage.Omnipage

func SetOmnipage(opInstance *omnipage.Omnipage) {
	op = opInstance
}

type request struct {
	Base64 string
}

type response struct {
	Text string
}

func ImgToText(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "application/json")

	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]

	if contentType == "application/json" {
		decoder := json.NewDecoder(r.Body)

		var jr request
		err := decoder.Decode(&jr)

		if err != nil {
			panic(err)
		}

		fmt.Println(jr.Base64)

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

			filename := "./aux_" + part.FileName()

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
		http.Error(w, errMsg, 400)

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
