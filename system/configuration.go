package system

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
)

// Configuration structure
type Configuration struct {
	Secret       string
	PublicPath   string
	TemplatePath string
	Database     struct {
		Hosts    string
		Database string
	}
}

// Load the configuration
func (configuration *Configuration) Load(filename string) (err error) {
	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(buf, &configuration); err != nil {
		panic(err)
	}

	return
}
