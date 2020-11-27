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
	UESubnet string    `yaml:"ueSubnet,omitempty"`
	UEList   *[]UE     `yaml:"ue,omitempty"`
	Snssai   *[]Snssai `yaml:"snssai,omitempty"`
	DNN      []string  `yaml:"dnn,omitempty"`
}

type UE struct {
	Supi string `yaml:"supi,omitempty"`
	PLMN string `yaml:"plmn,omitempty"`
}

type Snssai struct {
	Sst int32  `yaml:"sst,omitempty"`
	Sd  string `yaml:"sd,omitempty"`
}
