package qos

type Config struct {
	Info          *Info          `yaml:"info"`
	Configuration *Configuration `yaml:"configuration"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
}

type Configuration struct {
	Ipv4 []string `yaml:"ip,omitempty"`
	Port []string `yaml:"port,omitempty"`
}
