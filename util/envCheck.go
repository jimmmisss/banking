package util

import (
	"fmt"
	"github.com/jimmmisss/banking-lib/logger"
)

func EnvCheck() Config {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	envProps := map[string]string{
		"SERVER_ADDRESS": config.Server.Address,
		"SERVER_PORT":    config.Server.Port,
		"DB_USER":        config.DB.User,
		"DB_PASSWD":      config.DB.Password,
		"DB_ADDRESS":     config.DB.Address,
		"DB_PORT":        config.DB.Port,
		"DB_NAME":        config.DB.Name,
	}
	var envsErrs []string
	for i, k := range envProps {
		if k == "" {
			envsErrs = append(envsErrs, i)
		}
	}
	if len(envsErrs) != 0 {
		logger.Fatal(fmt.Sprintf("Env variable not defined %s. Finalizing application...", envsErrs))
	}
	return config
}
