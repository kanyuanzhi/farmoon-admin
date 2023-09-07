package config

type Config struct {
	System   System   `yaml:"system"`
	Postgres Postgres `yaml:"postgres"`
}

type System struct {
	Port           int    `yaml:"port"`
	SuccessCode    int    `yaml:"successCode"`
	SuccessMessage string `yaml:"successMessage"`
	ErrorCode      int    `yaml:"errorCode"`
	ErrorMessage   string `yaml:"errorMessage"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
