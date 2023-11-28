package config

import "github.com/spf13/viper"

//
// bilibiliVideo
//

type BilibiliVideoTarget struct {
	Uid      string `yaml:"uid"`
	Category string `yaml:"category"`
}

func GetBilibiliVideoSource() string {
	return viper.GetString("catcher.bilibiliVideo.source")
}

func GetBilibiliVideoTargets() []BilibiliVideoTarget {
	results := make([]BilibiliVideoTarget, 0)
	cfgs := viper.Get("catcher.bilibiliVideo.target").([]interface{})
	for _, c := range cfgs {
		results = append(results, BilibiliVideoTarget{
			Uid:      c.(map[string]interface{})["uid"].(string),
			Category: c.(map[string]interface{})["category"].(string),
		})
	}
	return results
}

//
// rssAdapter
//

type RssAdapterTarget struct {
	Url      string `yaml:"url"`
	Source   string `yaml:"source"`
	Category string `yaml:"category"`
}

func GetRssAdapterTargets() []RssAdapterTarget {
	results := make([]RssAdapterTarget, 0)
	cfgs := viper.Get("catcher.rssAdapter.target").([]interface{})
	for _, c := range cfgs {
		results = append(results, RssAdapterTarget{
			Url:      c.(map[string]interface{})["url"].(string),
			Source:   c.(map[string]interface{})["source"].(string),
			Category: c.(map[string]interface{})["category"].(string),
		})
	}
	return results
}
