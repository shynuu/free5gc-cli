package subscriber

type Config struct {
	Info          *Info          `yaml:"info"`
	Configuration *Configuration `yaml:"configuration"`
	PLMN          *PLMN          `yaml:"plmn"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
}

type Configuration struct {
	Mongodb *Mongodb `yaml:"mongodb"`
}

type Mongodb struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type PLMN struct {
	Plmn []string `yaml:"value"`
}
