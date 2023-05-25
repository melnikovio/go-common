package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/melnikovio/go-common/pkg/common"
)

func InitConfigurationOptions() {
	if v := os.Getenv(common.GetEnvConfigurationVariable()); v != "" {
		chantsEnabled, err := strconv.ParseBool(v)
		if err != nil {
			fmt.Println(fmt.Printf("Wrong %s initialization options", common.GetEnvConfigurationVariable()))
		}
		if chantsEnabled {
			params := common.GetConfigurationParameters()
			fmt.Println(common.GetConfigurationParametersDelimiter())
			for i := range params {
				fmt.Println(common.GetConfigurationParametersDelimiter())
				time.Sleep(500 * time.Millisecond)
				fmt.Println(params[i])
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Println(common.GetConfigurationParametersDelimiter())
		}
	} else {
		fmt.Println(
			fmt.Printf("Add custom init options passing %s parameter\n", common.GetEnvConfigurationVariable()))
	}
}
