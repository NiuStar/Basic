package Basic

import (
	"encoding/xml"
	"github.com/NiuStar/utils"
)

type Config struct {
	XMLName xml.Name 	`xml:"config"`
	LogConfig *logConfig `xml:"logpara"`
	DBConfig *DBConfig `xml:"db_info"`
	FileConfig *fileConfig `xml:"file"`
	ServerConfig *serverConfig `xml:"manage"`
}

type serverConfig struct {
	Listen_port string `xml:"listen_port"`
	TLS bool `xml:"tls"`
	Cert string `xml:"cert"`
	Key string `xml:"key"`
	Websocket bool `xml:"websocket"`
	Domain string `xml:"domain"`
}

func (s *serverConfig)GetPort() string {
	return s.Listen_port
}

type logConfig struct {
	Log_level string `xml:"log_level"`
	Log_path string `xml:"log_path"`
	Log_days int `xml:"log_days"`
}

type DBConfig struct {
	DB_server string `xml:"db_server"`
	DB_port string `xml:"db_port"`
	DB_name string `xml:"db_name"`
	DB_user string `xml:"db_user"`
	DB_password string `xml:"db_password"`
	DB_charset string `xml:"db_charset"`

}

func (db *DBConfig)GetDBSourceName() string {
	dataSourceName := db.DB_user + ":" + db.DB_password +
		"@tcp(" + db.DB_server + ":" + db.DB_port + ")/" + db.DB_name
	if len(db.DB_charset) > 0 {
		dataSourceName += "?charset=" + db.DB_charset
	}
	return dataSourceName
}

type fileConfig struct {
	LocalPath string
	NetworkPath string
	WebWorkDir string
}

var (
	gConfig *Config
)
//初始化;
func init() {
	gConfig = initConfig()
}

func initConfig() *Config {
	if utils.CheckFileIsExist("./Config/config.xml") {
		data := utils.ReadFileFullPath("./Config/config.xml")
		var config Config
		err := xml.Unmarshal([]byte(data),&config)
		if err != nil {
			panic(err)
		}
		return &config
	}
	return nil
}

func GetServerConfig() *Config {
	return gConfig
}