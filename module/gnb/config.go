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
	UESubnet     string       `yaml:"ueSubnet,omitempty"`
	UEList       *[]UE        `yaml:"ue,omitempty"`
	Snssai       *[]Snssai    `yaml:"snssai,omitempty"`
	DNN          []string     `yaml:"dnn,omitempty"`
	TUN          string       `yaml:"tun,omitempty"`
	GTPInterface GTPInterface `yaml:"gtpInterface,omitempty"`
}

type UE struct {
	Supi string `yaml:"supi,omitempty"`
	PLMN string `yaml:"plmn,omitempty"`
}

type Snssai struct {
	Sst int32  `yaml:"sst,omitempty"`
	Sd  string `yaml:"sd,omitempty"`
}

type GTPInterface struct {
	Ipv4 string `yaml:"ipv4Addr,omitempty"`
	Port int    `yaml:"port,omitempty"`
}
