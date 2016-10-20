package nuance

/*
#cgo CPPFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo LDFLAGS: -L /usr/local/lib/nuance-omnipage-csdk-lib64-19.2 -lrecapiplus -lkernelapi -lrecpdf -Wl,-rpath-link,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2,-rpath,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2

#include <KernelApi.h>

#include "nuancec.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type nuance struct {
	nuancePtr C.nuancePtr
}

func New() (n nuance) {
	n.nuancePtr = C.nuanceNew()
	return
}

func (n *nuance) Free() {
	C.nuanceFree(unsafe.Pointer(n.nuancePtr))
}

func (n *nuance) Init(company string, product string) (err error) {
	errBuff := make([]byte, 1024)
	if C.nuanceInit(
		unsafe.Pointer(n.nuancePtr),
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

func (n *nuance) SetLicense(licenseFile string, oemCode string) (err error) {
	errBuff := make([]byte, 1024)

	if C.nuanceSetLicense(
		unsafe.Pointer(n.nuancePtr),
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

func (n *nuance) Quit() {
	C.nuanceQuit(unsafe.Pointer(n.nuancePtr))
}

func (n *nuance) LoadFormTemplateLibrary(templateFile string) (err error) {
	errBuff := make([]byte, 1024)
	if C.nuanceLoadFormTemplateLibrary(
		unsafe.Pointer(n.nuancePtr),
		C.CString(templateFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func (n *nuance) OCRImgWithTemplate(imgFile string) (ret map[string]string, err error) {
	errBuff := make([]byte, 1024)
	ret = make(map[string]string)

	if C.nuancePreprocessImgWithTemplate(
		unsafe.Pointer(n.nuancePtr),
		C.CString(imgFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	zoneCount := int(C.nuanceGetZoneCount(unsafe.Pointer(n.nuancePtr)))

	fmt.Println("zoneCount:", zoneCount)

	for i := 0; i < zoneCount; i++ {
		zoneName := make([]byte, 256)
		zoneText := make([]byte, 256)

		C.nuanceGetZoneData(
			unsafe.Pointer(n.nuancePtr),
			C.int(i),
			(*C.char)(unsafe.Pointer(&zoneName[0])),
			C.int(256),
			(*C.char)(unsafe.Pointer(&zoneText[0])),
			C.int(256))

		ret[string(zoneName)] = string(zoneText)
		//fmt.Printf("%s: [%s]\n", string(zoneName), string(zoneText))
	}

	C.nuanceFreeImgWithTemplate(unsafe.Pointer(n.nuancePtr))
	err = nil
	return
}

func (n *nuance) OCRImgToText(imgFile string,
	outputFile string,
	auxDocumentFile string) (err error) {
	errBuff := make([]byte, 1024)

	if C.nuanceOCRImgToText(
		unsafe.Pointer(n.nuancePtr),
		C.CString(imgFile),
		C.CString(outputFile),
		C.CString(auxDocumentFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func (n *nuance) SetLanguagePtBr() (err error) {
	errBuff := make([]byte, 1024)

	if C.nuanceSetLanguagePtBr(
		unsafe.Pointer(n.nuancePtr),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}
