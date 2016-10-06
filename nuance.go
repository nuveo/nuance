package nuance

/*
// #cgo CFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo CPPFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo LDFLAGS: -L /usr/local/lib/nuance-omnipage-csdk-lib64-19.2 -lrecapiplus -lkernelapi -lrecpdf -Wl,-rpath-link,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2,-rpath,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2

#include <KernelApi.h>

#include "nuance.h"
*/
import "C"

func Quit() {
	C.Quit()
}

func SetLicense(licenseFile string, oemCode string) (r bool) {
	r = C.SetLicense(C.CString(licenseFile), C.CString(oemCode)) == 0
	return
}

func InitPDF(company string, product string) (r bool) {
	r = C.InitPDF(C.CString(company), C.CString(product)) == 0
	return
}

func LoadFormTemplateLibrary(templateFile string) (r bool) {
	r = C.LoadFormTemplateLibrary(C.CString(templateFile)) == 0
	return
}

func Nuance(filePath string) (res map[string]string, err error) {
	return
}
