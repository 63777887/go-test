package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	log_ "log"
	"os"
)

func init() {
	serverConfig, err := CreateServerConfig("./conf/config.yml")
	log_.Println("serverConfig: ", serverConfig)
	if err != nil {
		panic("配置文件初始化失败: " + err.Error())
	}
	Config = *serverConfig
}

var (
	Config ServerConfig
)

type ServerConfig struct {
	Server *server `json:"server,omitempty" yaml:"server,omitempty"`
	Mysql  *mysql  `json:"mysql,omitempty" yaml:"mysql,omitempty"`
	Log    *log    `json:"log,omitempty" yaml:"log,omitempty"`
}

type server struct {
	Host       string `json:"host,omitempty" yaml:"host,omitempty"`
	Port       int    `json:"port,omitempty" yaml:"port,omitempty"`
	Context    string `json:"context,omitempty" yaml:"context,omitempty"`
	Data       string `json:"data,omitempty" yaml:"data,omitempty"`
	BackupsDir string `json:"-" yaml:"-"`
	TempDir    string `json:"-" yaml:"-"`
}

type mysql struct {
	Host     string `json:"host,omitempty" yaml:"host,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Database string `json:"database,omitempty" yaml:"database,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

type log struct {
	Filename   string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MaxSize    int    `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
	MaxAge     int    `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
	MaxBackups int    `json:"maxBackups,omitempty" yaml:"maxBackups,omitempty"`
	Level      string `json:"level,omitempty" yaml:"level,omitempty"`
	Compress   bool   `json:"compress,omitempty" yaml:"compress,omitempty"`
}

func CreateServerConfig(configPath string) (config *ServerConfig, err error) {

	config = &ServerConfig{}
	if configPath != "" {
		var exists bool
		exists, err = PathExists(configPath)
		if err != nil {
			panic(err)
		}
		if !exists {
			err = errors.New(fmt.Sprint("服务配置文件[", configPath, "]不存在"))
			return
		}
		var f *os.File
		f, err = os.Open(configPath)
		if err != nil {
			panic(err)
		}
		yaml.NewDecoder(f).Decode(config)
	}

	if config.Server == nil {
		config.Server = &server{}
	}
	if config.Log == nil {
		config.Log = &log{
			MaxSize:    100,
			MaxAge:     7,
			MaxBackups: 10,
			Level:      "info",
		}
	}
	return
}

/*PathExists
  判断文件或文件夹是否存在
  如果返回的错误为nil,说明文件或文件夹存在
  如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
  如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
