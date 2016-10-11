package nuance

/*
// #cgo CFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo CPPFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo LDFLAGS: -L /usr/local/lib/nuance-omnipage-csdk-lib64-19.2 -lrecapiplus -lkernelapi -lrecpdf -Wl,-rpath-link,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2,-rpath,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2

#include <KernelApi.h>

#include "nuance.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func Quit() {
	C.Quit()
}

func SetLicense(licenseFile string, oemCode string) (err error) {
	errBuff := make([]byte, 1024)
	if C.SetLicense(C.CString(licenseFile),
		C.CString(oemCode),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func InitNuance(company string, product string) (err error) {
	errBuff := make([]byte, 1024)
	if C.InitNuance(
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

func LoadFormTemplateLibrary(templateFile string) (err error) {
	errBuff := make([]byte, 1024)
	if C.LoadFormTemplateLibrary(C.CString(templateFile),
		(*C.char)(unsafe.Pointer(&errBuff[0])),
		C.int(len(errBuff))) != 0 {

		err = errors.New(string(errBuff))
		return
	}

	err = nil
	return
}

func Nuance(filePath string) (res map[string]string, err error) {
	return
}
