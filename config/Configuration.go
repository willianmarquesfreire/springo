package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Cluster struct {
	Name       string   `yaml:"name"`
	DataCentre string   `yaml:"datacentre"`
	Nodes      []string `yaml:"nodes"`
}

type Configuration struct {
	Clusters    []Cluster `yaml:"clusters"`
	MinReplicas int       `yaml:"min_replicas"`
	MaxReplicas int       `yaml:"max_replicas"`
	Debug       bool      `yaml:"debug"`
	Database    string    `yaml:"database"`
	ApiPath     string    `yaml:"api_path"`
	Identifier  string    `yaml:"identifier"`
	MgoDial     string    `yaml:"mgo_dial"`
	Addr        string    `yaml:"addr"`
	BaseUrl     string    `yaml:"base_url"`
	SecurityUrl string    `yaml:"security_url"`
}

var MainConfiguration Configuration

func saveConfig(c Configuration, filename string) error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func loadConfig(filename string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}

func createMockConfigTest() Configuration {
	return Configuration{
		Clusters: []Cluster{
			Cluster{
				Name:       "Dev",
				DataCentre: "Local",
				Nodes:      []string{"dev1.company.com", "dev2.company.com"},
			},
			Cluster{
				Name:       "Prod",
				DataCentre: "Amazon",
				Nodes:      []string{"prd1.company.com", "prd2.company.com", "prd3.company.com"},
			},
		},
		MinReplicas: 1,
		MaxReplicas: 5,
		Debug:       true,
		Database:    "application_test",
		ApiPath:     "/application-api",
		Identifier:  "br.com.app",
		MgoDial:     "localhost",
		Addr:        ":8088",
		BaseUrl:     "/api/",
		SecurityUrl: "http://localhost:8088/application-api",
	}
}

func StartConfigurationTest(config Configuration) {
	log.Println("Starting configuration to test.")
	err := saveConfig(config, "config_test.yaml")
	if err != nil {
		panic(err)
	}

	c, err := loadConfig("config_test.yaml")
	if err != nil {
		panic(err)
	}

	MainConfiguration = c
}

func StartConfigurationTestMock() {
	StartConfigurationTest(createMockConfigTest())
}
