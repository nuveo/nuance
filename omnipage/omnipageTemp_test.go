package omnipage

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestAux1(t *testing.T) {
	oemCode := loadlicenseTxt()

	n := New()

	err := n.SetLicense(oemLicenseFile, oemCode)
	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	n.SetLanguagePtBr()

	err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
	if err != nil {
		t.Fatal("Init failed:", err)
	}

	p := "/src/aux/aux1.pdf"
	s := "/src/aux/aux1.txt"

	fmt.Println("convert:", p, s)

	err = n.OCRImgToTextFile(p, s, 0, "/src/aux/sample")
	if err != nil {
		//t.Fatal("OCRImgToTextFile failed:", err)
		fmt.Println("OCRImgToTextFile failed:", err)
	}

	n.Quit()
	n.Free()

}

func TestAux2(t *testing.T) {
	oemCode := loadlicenseTxt()

	for i := 1; i < 13; i++ {

		n := New()

		err := n.SetLicense(oemLicenseFile, oemCode)
		if err != nil {
			t.Fatal("SetLicense failed:", err)
		}

		err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
		if err != nil {
			t.Fatal("Init failed:", err)
		}

		p := fmt.Sprintf("/src/aux/aux%d.pdf", i)

		n.SetLanguagePtBr()

		x := 0
		x, err = n.CountPages(p)
		if err != nil {
			t.Fatal("CountPages failed:", err)
		}

		for y := 0; y < x; y++ {
			s := fmt.Sprintf("/src/aux/aux%d_p%d.txt", i, y+1)
			fmt.Println("\n\npagina", y+1, "do arquivo", p)

			err = n.OCRImgToTextFile(p, s, y, "/src/aux/sample")
			if err != nil {
				//t.Fatal("OCRImgToTextFile failed:", err)
				fmt.Println("failed:", err)
			}
		}

		n.Quit()
		n.Free()
	}
}

func TestAux3(t *testing.T) {
	oemCode := loadlicenseTxt()

	files, _ := ioutil.ReadDir("/src/aux/aux2/")

	for _, f := range files {
		//fmt.Println(f.Name())
		fName := filepath.Base(f.Name())
		extName := filepath.Ext(fName)
		bName := fName[:len(fName)-len(extName)]
		fmt.Println(bName, extName)

		n := New()

		err := n.SetLicense(oemLicenseFile, oemCode)
		if err != nil {
			t.Fatal("SetLicense failed:", err)
		}

		err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
		if err != nil {
			t.Fatal("Init failed:", err)
		}

		p := fmt.Sprintf("/src/aux/aux2/%s", fName)

		n.SetLanguagePtBr()

		x := 0
		x, err = n.CountPages(p)
		if err != nil {
			t.Fatal("CountPages failed:", err)
		}

		for y := 0; y < x; y++ {
			s := fmt.Sprintf("/src/aux/aux2/%s_p%d.txt", bName, y+1)
			fmt.Println("\n\npagina", y+1, "do arquivo", p)

			err = n.OCRImgToTextFile(p, s, y, "/src/aux/sample")
			if err != nil {
				//t.Fatal("OCRImgToTextFile failed:", err)
				fmt.Println("failed:", err)
			}
		}

		n.Quit()
		n.Free()
	}
	return

	for i := 1; i < 13; i++ {

		n := New()

		err := n.SetLicense(oemLicenseFile, oemCode)
		if err != nil {
			t.Fatal("SetLicense failed:", err)
		}

		err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
		if err != nil {
			t.Fatal("Init failed:", err)
		}

		p := fmt.Sprintf("/src/aux/aux%d.pdf", i)

		n.SetLanguagePtBr()

		x := 0
		x, err = n.CountPages(p)
		if err != nil {
			t.Fatal("CountPages failed:", err)
		}

		for y := 0; y < x; y++ {
			s := fmt.Sprintf("/src/aux/aux%d_p%d.txt", i, y+1)
			fmt.Println("\n\npagina", y+1, "do arquivo", p)

			err = n.OCRImgToTextFile(p, s, y, "/src/aux/sample")
			if err != nil {
				//t.Fatal("OCRImgToTextFile failed:", err)
				fmt.Println("failed:", err)
			}
		}

		n.Quit()
		n.Free()
	}
}

func TestAux4(t *testing.T) {
	oemCode := loadlicenseTxt()

	files, _ := ioutil.ReadDir("/src/aux/aux3/")

	for _, f := range files {
		//fmt.Println(f.Name())
		fName := filepath.Base(f.Name())
		extName := filepath.Ext(fName)
		bName := fName[:len(fName)-len(extName)]
		fmt.Println(bName, extName)

		n := New()

		err := n.SetLicense(oemLicenseFile, oemCode)
		if err != nil {
			t.Fatal("SetLicense failed:", err)
		}

		err = n.Init("YOUR_COMPANY", "YOUR_PRODUCT")
		if err != nil {
			t.Fatal("Init failed:", err)
		}

		p := fmt.Sprintf("/src/aux/aux3/%s", fName)

		n.SetLanguagePtBr()

		x := 0
		x, err = n.CountPages(p)
		if err != nil {
			t.Fatal("CountPages failed:", err)
		}

		for y := 0; y < x; y++ {
			s := fmt.Sprintf("/src/aux/aux3/%s_p%d.txt", bName, y+1)
			fmt.Println("\n\npagina", y+1, "do arquivo", p)

			err = n.OCRImgToTextFile(p, s, y, "/src/aux/sample")
			if err != nil {
				//t.Fatal("OCRImgToTextFile failed:", err)
				fmt.Println("failed:", err)
			}
		}

		n.Quit()
		n.Free()
	}
}

func TestNotasComTemplates(t *testing.T) {

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

	fmt.Println("LoadFormTemplateLibrary")
	err = n.LoadFormTemplateLibrary("/src/aux/nota3/template.ftl")
	if err != nil {
		t.Fatal("LoadFormTemplateLibrary failed:", err)
	}

	fmt.Println("OCRImgWithTemplate")
	var ret map[string]string
	ret, err = n.OCRImgWithTemplate("/src/aux/nota3/nota3.pdf")
	if err != nil {
		t.Fatal("OCRImgWithTemplate failed:", err)
	}

	for k, v := range ret {
		fmt.Println("k:", k, "v:", v)
	}

	n.Quit()
	n.Free()
}
