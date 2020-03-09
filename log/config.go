package log

import (
	"go.uber.org/zap/zapcore"
	"io/ioutil"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level zap.AtomicLevel `json:"level" yaml:"level"`
	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string `json:"encoding" yaml:"encoding"`
	//OutputPaths is a list of URLs or file paths to write logging output to.
	//See Open for details.
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
}

// init yaml config file
func InitWithFile(filePath string) (err error) {
	by, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	cf := &config{}
	err = yaml.Unmarshal(by, cf)
	if err != nil {
		return err
	}

	cnf := zap.NewProductionConfig()

	cnf.Level = cf.Level
	cnf.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cnf.DisableStacktrace = cf.DisableStacktrace

	switch cf.Encoding {
	case "console":
		cnf.Encoding = "console"
	case "json":
		cnf.Encoding = "json"
	default:
		cnf.Encoding = "json"
	}

	if len(cf.OutputPaths) > 0 {
		cnf.OutputPaths = cf.OutputPaths
	} else {
		cnf.OutputPaths = []string{"stdout"}
	}

	logger, err = cnf.Build()
	if err != nil {
		return err
	}

	return nil
}
