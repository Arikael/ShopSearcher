package search_config

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetAllConfigs(searchTerm string) (configs []*Config, err error) {
	var folder []os.DirEntry
	if folder, err = os.ReadDir(localFolderName); err != nil {
		return nil, err
	}

	for _, configFile := range folder {
		if config, err := GetConfig(configFile.Name(), searchTerm); err != nil {
			return nil, err
		} else {
			configs = append(configs, config)
		}
	}

	return configs, err
}

func GetConfig(configName string, searchTerm string) (config *Config, err error) {
	if !strings.HasSuffix(configName, ".txt") {
		configName += ".txt"
	}

	var (
		exists   bool
		filePath = filepath.Join(localFolderName, configName)
	)

	if exists, err = fsExists(filePath); err != nil {
		return nil, err
	} else if !exists {
		return nil, fmt.Errorf("config %s does not exist", filePath)
	}

	var file *os.File

	if file, err = os.Open(filePath); err != nil {
		return nil, err
	}

	defer func(file *os.File) error {
		err := file.Close()
		if err != nil {
			return err
		}

		return nil
	}(file)

	config = new(Config)
	config.Name = configName

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		var searchUrl = scanner.Text()
		if !strings.Contains(searchUrl, searchTermPattern) {
			continue
		}

		searchUrl = strings.Replace(searchUrl, searchTermPattern, url.QueryEscape(searchTerm), 1)
		config.Urls = append(config.Urls, searchUrl)
	}

	return config, nil
}
