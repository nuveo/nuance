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

func loadlicenseTxt() string {
	code, err := ioutil.ReadFile(oemLicenseFile)
	if err != nil {
		log.Fatal("Error loading license file", oemLicenseFile, err)
	}

	return string(code)
}

func TestSetLicense(t *testing.T) {

	oemCode := loadlicenseTxt()

	r := SetLicense(oemLicenseFile, oemCode)

	if !r {
		t.Fatal("SetLicense failed")
	}

	Quit()
}

func TestInitPDF(t *testing.T) {

	oemCode := loadlicenseTxt()

	r := SetLicense(oemLicenseFile, oemCode)

	if !r {
		t.Fatal("SetLicense failed")
	}

	r = InitPDF("YOUR_COMPANY", "YOUR_PRODUCT")

	if !r {
		t.Fatal("InitPDF failed")
	}

	Quit()
}

func TestLoadFormTemplateLibrary(t *testing.T) {

	oemCode := loadlicenseTxt()

	r := SetLicense(oemLicenseFile, oemCode)

	if !r {
		t.Fatal("SetLicense failed")
	}

	r = LoadFormTemplateLibrary("/src/template.ftl")

	if !r {
		t.Fatal("LoadFormTemplateLibrary failed")
	}

	Quit()
}
