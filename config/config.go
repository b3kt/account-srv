package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// ConfigFile is the default config file
var ConfigFile = "/home/bekt/go/src/github.com/b3kt/account-srv/config.yml"

// GlobalConfig is the global config
type GlobalConfig struct {
	Server        ServerConfig        `yaml:"server"`
	KeycloakAdmin KeycloakAdminConfig `yaml:"keycloak"`
	Redis         RedisConfig         `yaml:"redis"`
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
	BaseURL      string `yaml:"base_url"`
	AdminRealm   string `yaml:"admin_realm"`
	Realm        string `yaml:"realm"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// RedisConfig config integration to redis
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Expire   time.Duration
}

// global configs
var (
	Global        GlobalConfig
	Server        ServerConfig
	KeycloakAdmin KeycloakAdminConfig
	Redis         RedisConfig
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
	Redis = Global.Redis

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
