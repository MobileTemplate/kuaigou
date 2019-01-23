
package config

import (
	"bytes"
	"io/ioutil"
	"os"

	"qianuuu.com/lib/logs"

	"github.com/BurntSushi/toml"
)

// Config 配置类型
type Config struct {
	Port             string // 服务器端口
	ConnString       string // 数据库连接地址
	LogPath          string // 日志目录
}


// Opts 配置默认值
var Opts = Config{
	Port:                  "8106",
	ConnString:            "postgres://postgres:postgres@localhost:5432/poker?sslmode=disable",
	LogPath:               "",
}

// ParseToml 解析配置文件
func ParseToml(file string) error {
	logs.Info("读取配置文件 ...")
	// 如果配置文件不存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(Opts); err != nil {
			return err
		}
		logs.Info("没有找到配置文件，创建新文件 ...")
		// logs.Info(buf.String())
		return ioutil.WriteFile(file, buf.Bytes(), 0644)
	}
	var conf Config
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		return err
	}
	Opts = conf
	return nil
}
