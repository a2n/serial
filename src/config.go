package serial

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

type ConfigService struct{}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

type Config struct {
	Value uint64 `toml:"value"`
	Port  string `toml:"port"`
}

func (this *ConfigService) Get() (*Config, error) {
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

func (this *ConfigService) Save(c *Config) error {
	if c == nil {
		return errors.New("Nil config.")
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
