package controllers

import (
	"strings"
	"testing"

	"github.com/caarlos0/env"
	"github.com/nuveo/nuance/config"
	"github.com/nuveo/nuance/omnipage"
)

func TestOcrFile(t *testing.T) {
	cfg := config.Nuance{}
	env.Parse(&cfg)

	op := omnipage.New()

	err := op.SetLicense(cfg.OemLicenseFile, cfg.OemCode)
	if err != nil {
		t.Fatal("SetLicense failed:", err)
	}

	err = op.Init(cfg.CompanyName, cfg.ProductName)
	if err != nil {
		t.Fatal("Init failed:", err)
	}

	SetOmnipage(&op)
	SetConfig(&cfg)

	txt := ""
	txt, err = ocrFile("../test_assets/testFile.png")
	if err != nil {
		t.Fatal("ocrFile:", err)
	}
	txt = strings.Replace(txt, "  ", " ", -1)
	if !strings.Contains(txt, "It is a test.") {
		t.Error("Error: string does not contain expected substring.")
	}
}
