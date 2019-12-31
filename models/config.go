package models

type Config struct {
	Chartmuseum Chartmuseum `yaml:"chartmuseum"`
	UI          UI          `yaml:"ui"`
}

type Chartmuseum struct {
	Host     string `yaml:"host"`
	HostAPI  string `yaml:"hostapi"`
	Username string `yaml:"username",omitempty`
	Password string `yaml:"password",omitempty`
}

type UI struct {
	Username string `yaml:"username",omitempty`
	Password string `yaml:"password",omitemtpy`
}
