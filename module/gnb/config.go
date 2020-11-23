package gnb

type Config struct {
	Info          *Info          `yaml:"info,omitempty"`
	Configuration *Configuration `yaml:"configuration,omitempty"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
}

type Configuration struct {
	UESubnet       string         `yaml:"ueSubnet,omitempty"`
	UEList         *[]UE          `yaml:"ue,omitempty"`
	AmfInterface   AmfInterface   `yaml:"amfInterface,omitempty"`
	UpfInterface   UpfInterface   `yaml:"upfInterface,omitempty"`
	NGRANInterface NGRANInterface `yaml:"ngranInterface,omitempty"`
	GTPInterface   GTPInterface   `yaml:"gtpInterface,omitempty"`
	Security       Security       `yaml:"security,omitempty"`
	Snssai         *[]Snssai      `yaml:"snssai,omitempty"`
}

type UE struct {
	Supi string `yaml:"supi,omitempty"`
	PLMN string `yaml:"plmn,omitempty"`
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

type Snssai struct {
	Sst int32  `yaml:"sst,omitempty"`
	Sd  string `yaml:"sd,omitempty"`
}
