package config

// Nuance basic config
type Nuance struct {
	// HTTPPort Declare which http port the PREST used
	HTTPPort       int    `env:"NUANCE_HTTP_PORT" envDefault:"4000"`
	JWTKey         string `env:"NUANCE_JWT_KEY"`
	OemLicenseFile string `env:"NUANCE_OEM_LICENCE" envDefault:"/src/license.lcxz"`
	OemCode        string `env:"NUANCE_OEM_CODE" envDefault:"OEM_CODE"`
	CompanyName    string `env:"NUANCE_COMPANY_NAME" envDefault:"YOUR_COMPANY"`
	ProductName    string `env:"NUANCE_PRODUCT_NAME" envDefault:"YOUR_PRODUCT"`
	TmpPath        string `env:"NUANCE_TEMP_PATH" envDefault:"/tmp"`
}
