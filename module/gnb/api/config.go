package api

type Config struct {
	Configuration *Configuration `yaml:"configuration,omitempty"`
}

type Configuration struct {
	AmfInterface   AmfInterface   `yaml:"amfInterface,omitempty"`
	UpfInterface   UpfInterface   `yaml:"upfInterface,omitempty"`
	NGRANInterface NGRANInterface `yaml:"ngranInterface,omitempty"`
	GTPInterface   GTPInterface   `yaml:"gtpInterface,omitempty"`
	Security       Security       `yaml:"security,omitempty"`
}

type AmfInterface struct {
	IPv4Addr string `yaml:"ipv4Addr,omitempty"`
	Port     int    `yaml:"port,omitempty"`
}

type UpfInterface struct {
	IPv4Addr string `yaml:"ipv4Addr,omitempty"`
	Port     int    `yaml:"port,omitempty"`
}

type NGRANInterface struct {
	IPv4Addr string `yaml:"ipv4Addr,omitempty"`
	Port     int    `yaml:"port,omitempty"`
}

type GTPInterface struct {
	IPv4Addr string `yaml:"ipv4Addr,omitempty"`
	Port     int    `yaml:"port,omitempty"`
}

type Security struct {
	NetworkName string `yaml:"networkName,omitempty"`
	K           string `yaml:"k,omitempty"`
	OPC         string `yaml:"opc,omitempty"`
	OP          string `yaml:"op,omitempty"`
	SQN         string `yaml:"sqn,omitempty"`
}
