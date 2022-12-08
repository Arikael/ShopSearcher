package main

import (
	"flag"
	"github.com/Arikael/shop-searcher/browser"
	"github.com/Arikael/shop-searcher/search_config"
	"log"
	"strings"
)

func main() {
	var err error
	if err = search_config.InitLocalData(); err != nil {
		log.Fatal(err)
	}

	configName := flag.String("config", "", "config file to use (found in ./data folder)")

	flag.Parse()
	var (
		searchArgs = flag.Args()
		configs    []*search_config.Config
	)

	if len(*configName) == 0 {
		if configs, err = search_config.GetAllConfigs(strings.Join(searchArgs, " ")); err != nil {
			log.Fatal(err)
		}
	} else {
		var config *search_config.Config
		if config, err = search_config.GetConfig(*configName, strings.Join(searchArgs, " ")); err != nil {
			log.Fatal(err)
		} else {
			configs = append(configs, config)
		}
	}

	for _, config := range configs {
		if err = browser.OpenBrowser(config); err != nil {
			log.Fatal(err)
		}
	}
}
