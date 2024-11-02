package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
	FilePath string
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) LoadEnv() {
	start, err := parseStartArgument()
	if err != nil {
		log.Println(ErrEnvParseStart, err.Error())
		return
	}

	if err := loadConfigFile(*start); err != nil {
		log.Println(ErrEnvLoad, err.Error())
		return
	}

	e.FilePath = *start
	setEnvironmentVariables()
}

func parseStartArgument() (*string, error) {
	for key, value := range os.Args {
		if key == 0 {
			continue
		}

		t := strings.Split(value, "=")
		if len(t) != 2 {
			continue
		}

		m := strings.ToUpper(strings.Trim(t[0], "- "))
		if m == "START" {
			start := t[1]
			os.Setenv("API_DIRENV", start)
			return &start, nil
		}
	}
	return nil, fmt.Errorf(ErrEnvNoStart)
}
func loadConfigFile(start string) error {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(start)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf(ErrEnvNoConfigFile, start)
	}

	return nil
}
func setEnvironmentVariables() {
	var paramStart string

	for key, value := range os.Args {
		if key == 0 {
			continue
		}

		t := strings.Split(value, "=")
		if len(t) != 2 {
			continue
		}

		m := strings.ToUpper(strings.Trim(t[0], "- "))
		os.Setenv(m, t[1])
		paramStart += fmt.Sprintf("-%s=%s ", m, t[1])
	}

	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		os.Setenv(envKey, value)
		paramStart += fmt.Sprintf("-%s=%s ", key, value)
	}

	os.Setenv("API_START", paramStart)
}
