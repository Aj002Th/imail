package config

import "github.com/spf13/viper"

//
// bilibiliVideo
//

type BilibiliVideoConfig struct {
	Uid      string `yaml:"uid"`
	Category string `yaml:"category"`
}

func GetBilibiliVideoSource() string {
	return viper.GetString("catcher.bilibiliVideo.source")
}

func GetBilibiliVideoConfigs() []BilibiliVideoConfig {
	results := make([]BilibiliVideoConfig, 0)
	cfgs := viper.Get("catcher.bilibiliVideo.target").([]interface{})
	for _, c := range cfgs {
		results = append(results, BilibiliVideoConfig{
			Uid:      c.(map[string]interface{})["uid"].(string),
			Category: c.(map[string]interface{})["category"].(string),
		})
	}
	return results
}
