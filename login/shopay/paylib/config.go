//
// Author: leafsoar
// Date: 2018-06-21 16:47:20
//

package paylib

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"qianuuu.com/lib/logs"
)

// Config 主配置
type Config struct {
	Title   string // 标题
	AppName string
	RootURL string // 通知地址 根地址
	Items   []PayItem
}

var opts = Config{
	Title:   "支付",
	RootURL: "http://bing.com",
	Items: []PayItem{
		PayItem{
			AppName:  "name",
			PayName:  "test",
			PayTypes: []string{"wxh5", "alih5"},
		},
	},
}

func Opts() *Config {
	return &opts
}

func (c *Config) GetItem(payname string) (*PayItem, error) {
	for _, item := range c.Items {
		logs.Info("--->", item)
		if item.PayName == payname {
			return &item, nil
		}
	}
	return nil, errors.New("没有找到支付通道")
}

func ParseToml(file string) error {
	logs.Info("读取配置文件 ...")
	// 如果配置文件不存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(opts); err != nil {
			return err
		}
		logs.Info("%v", opts)
		logs.Info("没有找到配置文件，创建新文件 ...")
		return ioutil.WriteFile(file, buf.Bytes(), 0644)
	}
	var conf Config
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		return err
	}
	opts = conf
	return nil
}
