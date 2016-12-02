package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

    "github.com/nuveo/nuance/config"
    "github.com/nuveo/nuance/omnipage"
)

func PostImg(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	//TODO: Check file MIME type

	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

    // -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

    n := omnipage.New()

    err := n.SetLicense(oemLicenseFile, oemCode)
    if err != nil {
        t.Fatal("SetLicense failed:", err)
    }

    err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
    if err != nil {
        t.Fatal("Init failed:", err)
    }

    n.SetLanguagePtBr()

    n.SetCodePage("UTF-8")
    if err != nil {
        t.Fatal("SetCodePage failed:", err)
    }

    var txt string
    txt, err = n.OCRImgToText("/src/sample.tif")
    if err != nil {
        t.Fatal("TestCodePage failed:", err)
    }

    fmt.Println("TestCodePage:", txt)

    // REMOVE file

}
