package nuance

/*
// #cgo CFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo CPPFLAGS: -I /usr/local/include/nuance-omnipage-csdk-19.2
#cgo LDFLAGS: -L /usr/local/lib/nuance-omnipage-csdk-lib64-19.2 -lrecapiplus -lkernelapi -lrecpdf -Wl,-rpath-link,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2,-rpath,/usr/local/lib/nuance-omnipage-csdk-lib64-19.2

#include "nuance.h"
*/
import "C"

func SetLicense(licenceFile string, oemCode string) (r bool) {
	r = C.SetLicense(C.CString(licenceFile), C.CString(oemCode)) == 0
	return
}

func Nuance(filePath string) (res map[string]string, err error) {
	return
}
