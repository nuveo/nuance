package nuance

import (
	"fmt"
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

	n := New()
	err := n.SetLicense(oemLicenseFile, oemCode)

	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	n.Quit()
	n.Free()
}

func TestInit(t *testing.T) {

	oemCode := loadlicenseTxt()

	n := New()

	err := n.SetLicense(oemLicenseFile, oemCode)
	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
	if err != nil {
		t.Fatal("Init failed:", err)
	}

	n.Quit()
	n.Free()

}

func TestLoadFormTemplateLibrary(t *testing.T) {

	oemCode := loadlicenseTxt()

	n := New()

	err := n.SetLicense(oemLicenseFile, oemCode)
	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = n.LoadFormTemplateLibrary("/src/template.ftl")
	if err != nil {
		t.Fatal("LoadFormTemplateLibrary failed:", err)
	}

	n.Quit()
	n.Free()

}

func TestOCRImgWithTemplate(t *testing.T) {

	oemCode := loadlicenseTxt()

	n := New()

	err := n.SetLicense(oemLicenseFile, oemCode)
	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
	if err != nil {
		t.Fatal("Init failed:", err)
	}

	err = n.LoadFormTemplateLibrary("/src/template.ftl")
	if err != nil {
		t.Fatal("LoadFormTemplateLibrary failed:", err)
	}

	var ret map[string]string
	ret, err = n.OCRImgWithTemplate("/src/sample.tif")
	if err != nil {
		t.Fatal("OCRImgWithTemplate failed:", err)
	}

	for k, v := range ret {
		fmt.Println("k:", k, "v:", v)
	}

	n.Quit()
	n.Free()
}

func TestOCRImgToFile(t *testing.T) {
	oemCode := loadlicenseTxt()

	n := New()

	err := n.SetLicense(oemLicenseFile, oemCode)
	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
	if err != nil {
		t.Fatal("Init failed:", err)
	}

	err = n.OCRImgToFile("/src/sample.tif",
		"/src/sample.txt",
		0,
		"/src/sample.doc")
	if err != nil {
		t.Fatal("OCRImgToText failed:", err)
	}

}

func TestOCRImgToText(t *testing.T) {
        oemCode := loadlicenseTxt()

        n := New()

        err := n.SetLicense(oemLicenseFile, oemCode)
        if err != nil {
                t.Fatal("SetLicense failed:", err)
        }

        err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
        if err != nil {
                t.Fatal("Init failed:", err)
        }

        txt, err := n.OCRImgToText("/src/sample.tif", 0)
        if err != nil {
                t.Errorf("Expected no errors, but foud %s", err)
        }

}

func TestMultiplePagesOCRImgToFile(t *testing.T) {
	oemCode := loadlicenseTxt()

	n := New()

	err := n.SetLicense(oemLicenseFile, oemCode)
	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
	if err != nil {
		t.Fatal("Init failed:", err)
	}

	err = n.OCRImgToFile("/src/sample.tif",
		"/src/sample.txt",
		0,
		"/src/sample.doc")
	if err != nil {
		t.Fatal("OCRImgToText failed:", err)
	}
}
