package config

type Config struct {
	System   System   `yaml:"system"`
	Postgres Postgres `yaml:"postgres"`
	Zap      Zap      `yaml:"zap"`
	RPC      RPC      `yaml:"rpc"`
}

type System struct {
	Port                          int    `yaml:"port"`
	GinReleaseMode                bool   `yaml:"ginReleaseMode"`
	SuccessCode                   int    `yaml:"successCode"`
	SuccessMessage                string `yaml:"successMessage"`
	ErrorCode                     int    `yaml:"errorCode"`
	ErrorMessage                  string `yaml:"errorMessage"`
	BindErrorMessage              string `yaml:"bindErrorMessage"`
	LoadLocalResourceErrorMessage string `yaml:"loadLocalResourceErrorMessage"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Zap struct {
	Prefix      string `yaml:"prefix"`
	Level       string `yaml:"level"`
	Path        string `yaml:"path"`
	Filename    string `yaml:"filename"`
	MaxSize     int    `yaml:"maxSize"`
	MaxBackups  int    `yaml:"maxBackups"`
	MaxAge      int    `yaml:"maxAge"`
	OsStdout    bool   `yaml:"osStdout"`
	JsonEncoder bool   `yaml:"jsonEncoder"`
}

type RPC struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
