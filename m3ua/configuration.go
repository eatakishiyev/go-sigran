package m3ua

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

type M3UAConfiguration struct {
	AsSelectMask                   uint32 `mapstructure:"as-select-mask"`
	AspSelectMask                  uint32 `mapstructure:"asp-select-mask"`
	MaxAsCountForRouting           uint32 `mapstructure:"max-as-count-for-routing"`
	MaxSls                         uint16 `mapstructure:"max-sls"`
	DelayBetweenInitiationMessages uint32 `mapstructure:"delay-between-initiation-messages"`
}

var M3UAConfig M3UAConfiguration

func init() {
	var config = viper.New()
	config.SetConfigFile("m3ua.json")
	config.ReadInConfig()

	err := config.Unmarshal(&M3UAConfig)
	if err != nil {
		fmt.Printf("error occurred during read m3ua configuration %s\n", err)
	}
}
