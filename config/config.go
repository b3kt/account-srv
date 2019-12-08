package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// ConfigFile is the default config file
var ConfigFile = "./config.yml"

// GlobalConfig is the global config
type GlobalConfig struct {
	Server        ServerConfig        `yaml:"server"`
	KeycloakAdmin KeycloakAdminConfig `yaml:"keycloak"`
}

// ServerConfig is the server config
type ServerConfig struct {
	Addr               string
	Mode               string
	Version            string
	StaticDir          string `yaml:"static_dir"`
	ViewDir            string `yaml:"view_dir"`
	LogDir             string `yaml:"log_dir"`
	UploadDir          string `yaml:"upload_dir"`
	MaxMultipartMemory int64  `yaml:"max_multipart_memory"`
}

// KeycloakAdminConfig keycloak admin config
type KeycloakAdminConfig struct {
	BaseURL    string `yaml:"base_url"`
	AdminRealm string `yaml:"adminRealm"`
	Realm      string `yaml:"realm"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

// global configs
var (
	Global        GlobalConfig
	Server        ServerConfig
	KeycloakAdmin KeycloakAdminConfig
)

// Load config from file
func Load(file string) (GlobalConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	err = yaml.Unmarshal(data, &Global)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	Server = Global.Server
	KeycloakAdmin = Global.KeycloakAdmin

	// set log dir flag for glog
	flag.CommandLine.Set("log_dir", Server.LogDir)

	return Global, nil
}

// loads configs
func init() {
	if os.Getenv("config") != "" {
		ConfigFile = os.Getenv("config")
	}
	Load(ConfigFile)
}
