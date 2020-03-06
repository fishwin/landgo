package log

import (
	"io/ioutil"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type config struct {
	Level zap.AtomicLevel `json:"level"`
	// Encoding json or console
	Encoding string `json:"encoding,omitempty"`
	// Outputs 输出目录，支持文件，stdout stderr 等
	Outputs []string `json:"outputs,omitempty"`
}

// InitWithFile init yaml config file
func InitWithFile(filePath string) (err error) {
	by, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	cf := config{}
	err = yaml.Unmarshal(by, cf)
	if err != nil {
		return err
	}

	return nil
}
