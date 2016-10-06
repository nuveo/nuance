package nuance

import (
	"io/ioutil"
	"log"
	"testing"
)

// place your OEM code in a plain text file
const oemLicenseTxtFile = "/src/license.txt"

// replace it with your license file
const oemLicenseFile = "/src/license.lcxz"

func loadLicenceTxt() string {
	code, err := ioutil.ReadFile(oemLicenseFile)
	if err != nil {
		log.Fatal("Error loading licence file", oemLicenseFile, err)
	}

	return string(code)
}

func TestSetLicense(t *testing.T) {

	oemCode := loadLicenceTxt()

	r := SetLicense(oemLicenseFile, oemCode)

	if !r {
		t.Fatal("SetLicense failed")
	}
}
