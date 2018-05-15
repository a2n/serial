package serial

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// ConfigService 配置服務
type ConfigService struct{}

// NewConfigService 創建服務
func NewConfigService() *ConfigService {
	return &ConfigService{}
}

// Config 配置
type Config struct {
	Value uint64 `toml:"value"`
	Port  string `toml:"port"`
}

// Get 取得
func (cs *ConfigService) Get() (*Config, error) {
	b, e := ioutil.ReadFile("config/serial.toml")
	if e != nil {
		return nil, errors.Wrap(e, "")
	}

	c := &Config{}
	e = toml.Unmarshal(b, &c)
	if e != nil {
		return nil, errors.Wrap(e, "")
	}

	return c, nil
}

// Save 保存
func (cs *ConfigService) Save(c *Config) error {
	if c == nil {
		return errors.New("nil config")
	}

	// Open file.
	f, e := os.OpenFile("config/serial.toml", os.O_RDWR, 0600)
	if e != nil {
		return errors.Wrap(e, "")
	}
	defer func() {
		e = f.Close()
		if e != nil {
			glog.Errorf("%+v", e)
			return
		}
	}()

	// Encoding.
	e = toml.NewEncoder(f).Encode(c)
	if e != nil {
		return errors.Wrap(e, "")
	}

	return nil
}
