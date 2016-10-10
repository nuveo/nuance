package nuance

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

// place your OEM code in a plain text file
const oemLicenseTxtFile = "/src/license.txt"

// replace it with your license file
const oemLicenseFile = "/src/license.lcxz"

func loadlicenseTxt() (r string) {
	code, err := ioutil.ReadFile(oemLicenseTxtFile)
	if err != nil {
		log.Fatal("Error loading license file", oemLicenseTxtFile, err)
	}

	r = strings.TrimSpace(string(code))

	return r
}

func TestSetLicense(t *testing.T) {

	oemCode := loadlicenseTxt()

	err := SetLicense(oemLicenseFile, oemCode)

	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	Quit()
}

func TestInitPDF(t *testing.T) {

	oemCode := loadlicenseTxt()

	err := SetLicense(oemLicenseFile, oemCode)

	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = InitPDF("YOUR_COMPANY", "YOUR_PRODUCT")

	if err != nil {
		t.Fatal("InitPDF failed:", err)
	}

	Quit()
}

func TestLoadFormTemplateLibrary(t *testing.T) {

	oemCode := loadlicenseTxt()

	err := SetLicense(oemLicenseFile, oemCode)

	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = LoadFormTemplateLibrary("/src/template.ftl")

	if err != nil {
		t.Fatal("LoadFormTemplateLibrary failed:", err)
	}

	Quit()
}
