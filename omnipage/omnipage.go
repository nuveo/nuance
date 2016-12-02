package omnipage

/*
#cgo CPPFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo LDFLAGS: -L /usr/local/lib/nuance-omnipage-csdk-lib64-19.2 -lrecapiplus -lkernelapi -lrecpdf -Wl,-rpath-link,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2,-rpath,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2

#include <KernelApi.h>

#include "omnipagec.h"
*/
import "C"

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"time"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

type omnipage struct {
	omnipagePtr C.omnipagePtr
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New create omnipage session
func New() (n omnipage) {
	n.omnipagePtr = C.omnipageNew()
	return
}

// Free allocated resources
func (n *omnipage) Free() {
	C.omnipageFree(unsafe.Pointer(n.omnipagePtr))
}

// Init a omnipage session.
func (n *omnipage) Init(company string, product string) (err error) {
	errBuff := make([]byte, 1024)
	if C.omnipageInit(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(company),
		C.CString(product),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

/*
SetLicense uses license file and OEM code to activate omnipage library.
It should be called immediately after instantiating a new session with the New () function.
*/
func (n *omnipage) SetLicense(licenseFile string, oemCode string) (err error) {
	errBuff := make([]byte, 1024)

	if C.omnipageSetLicense(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(licenseFile),
		C.CString(oemCode),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

/*
Quit omnipage.
At this point of development it is also necessary to use the
Free() function to release the resources.
*/
func (n *omnipage) Quit() {
	C.omnipageQuit(unsafe.Pointer(n.omnipagePtr))
}

func (n *omnipage) LoadFormTemplateLibrary(templateFile string) (err error) {
	errBuff := make([]byte, 1024)
	if C.omnipageLoadFormTemplateLibrary(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(templateFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func (n *omnipage) OCRImgWithTemplate(imgFile string) (ret map[string]string, err error) {
	errBuff := make([]byte, 1024)
	ret = make(map[string]string)

	if C.omnipagePreprocessImgWithTemplate(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(imgFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	zoneCount := int(C.omnipageGetZoneCount(unsafe.Pointer(n.omnipagePtr)))

	fmt.Println("zoneCount:", zoneCount)

	for i := 0; i < zoneCount; i++ {
		zoneName := make([]byte, 256)
		zoneText := make([]byte, 256)

		C.omnipageGetZoneData(
			unsafe.Pointer(n.omnipagePtr),
			C.int(i),
			(*C.char)(unsafe.Pointer(&zoneName[0])),
			C.int(256),
			(*C.char)(unsafe.Pointer(&zoneText[0])),
			C.int(256))

		ret[string(zoneName)] = string(zoneText)
		//fmt.Printf("%s: [%s]\n", string(zoneName), string(zoneText))
	}

	C.omnipageFreeImgWithTemplate(unsafe.Pointer(n.omnipagePtr))
	err = nil
	return
}

func (n *omnipage) OCRImgToFile(imgFile string,
	outputFile string,
	nPage int,
	auxDocumentFile string) (err error) {
	errBuff := make([]byte, 1024)

	if C.omnipageOCRImgToFile(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(imgFile),
		C.CString(outputFile),
		C.int(nPage),
		C.CString(auxDocumentFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}
	return
}

func (n *omnipage) OCRImgToTextFile(imgFile string,
	outputFile string,
	nPage int,
	auxDocumentFile string) (err error) {
	errBuff := make([]byte, 1024)

	randomAux := randString(6)
	tempFile := fmt.Sprintf("%s.%s", outputFile, randomAux)

	defer func() {
		os.Remove(tempFile)
	}()

	if C.omnipageOCRImgToTextFile(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(imgFile),
		C.CString(tempFile),
		C.int(nPage),
		C.CString(auxDocumentFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	var iArray []byte
	iArray, err = ioutil.ReadFile(tempFile)
	if err != nil {
		fmt.Println("OCRImgToTextFile error:", err)
		return
	}

	l := len(iArray)
	if l%2 != 0 {
		err = errors.New("OCRImgToTextFile Number of bytes in the file must be multiple of 2")
		return
	}

	u16s := make([]uint16, 1)
	b8buf := make([]byte, 4)
	oArray := &bytes.Buffer{}

	for i := 0; i < l; i += 2 {
		u16s[0] = uint16(iArray[i]) + (uint16(iArray[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		oArray.Write(b8buf[:n])
	}

	err = ioutil.WriteFile(outputFile, oArray.Bytes(), 0644)
	if err != nil {
		fmt.Println("OCRImgToTextFile error:", err)
		return
	}

	err = nil
	return
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (n *omnipage) OCRImgPageToText(imgFile string, nPage int) (txt string, err error) {
	randomAux := randString(6)
	tempDir := path.Join(os.TempDir(), randomAux)
	tempFile := fmt.Sprintf("%s.txt", tempDir)

	defer func() {
		os.Remove(tempFile)
		os.RemoveAll(tempDir)
	}()
	err = n.OCRImgToTextFile(imgFile, tempFile, nPage, tempDir)
	if err != nil {
		fmt.Println("OCRImgPageToText error:", err)
		return
	}
	rawTxt, err := ioutil.ReadFile(tempFile)
	if err != nil {
		fmt.Println("OCRImgPageToText error:", err)
		return
	}
	txt = string(rawTxt)

	return
}

func (n *omnipage) OCRImgToText(imgFile string) (txt string, err error) {
	var pages int
	pages, err = n.CountPages(imgFile)
	if err != nil {
		return
	}
	var aux string
	for i := 0; i < pages; i++ {
		aux, err = n.OCRImgPageToText(imgFile, i)
		if err != nil {
			fmt.Println("OCRImgToText error: ", err)
			return
		}
		if len(txt) > 0 {
			txt += "\n"
		}
		txt += aux
	}
	return
}

func (n *omnipage) SetLanguagePtBr() (err error) {
	errBuff := make([]byte, 1024)

	if C.omnipageSetLanguagePtBr(
		unsafe.Pointer(n.omnipagePtr),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func (n *omnipage) CountPages(imgFile string) (nPage int, err error) {
	errBuff := make([]byte, 1024)
	nPage = 0

	if C.omnipageCountPages(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(imgFile),
		(*C.int)(unsafe.Pointer(&nPage)),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func (n *omnipage) SetCodePage(codePage string) (err error) {
	errBuff := make([]byte, 1024)

	if C.omnipageSetCodePage(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(codePage),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func (n *omnipage) SetOutputFormat(outputFormat string) (err error) {
	errBuff := make([]byte, 1024)

	if C.omnipageSetOutputFormat(
		unsafe.Pointer(n.omnipagePtr),
		C.CString(outputFormat),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}
