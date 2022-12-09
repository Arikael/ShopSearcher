package search_config

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func InitLocalData() (err error) {
	var localFolder string

	if localFolder, err = createLocalDataDir(); err != nil {
		return err
	}

	files, err := embeddedDataFolder.ReadDir(embeddedFolderName)

	for _, file := range files {
		if err = createLocalFile(localFolder, file); err != nil {
			return err
		}
	}

	return nil
}

func createLocalDataDir() (localFolder string, err error) {
	var wd string
	wd, err = os.Getwd()
	var folderExists bool
	localFolder = filepath.Join(wd, localFolderName)
	folderExists, err = fsExists(localFolder)

	if err == nil && !folderExists {
		err = os.Mkdir(localFolder, 0644)
		return "", err
	}

	return localFolder, nil
}

func createLocalFile(localFolder string, file fs.DirEntry) (err error) {
	var localFilePath = filepath.Join(localFolder, file.Name())
	// embedded paths must be with slash
	var embeddedFilePath = strings.Replace(filepath.Join(embeddedFolderName, file.Name()), "\\", "/", -1)
	var fileContent []byte

	if localFileExists, err := fsExists(localFilePath); err != nil {
		return err
	} else if !localFileExists {
		fileContent, err = embeddedDataFolder.ReadFile(file.Name())
		if fileContent, err = embeddedDataFolder.ReadFile(embeddedFilePath); err != nil {
			return err
		}

		if err = os.WriteFile(localFilePath, fileContent, 0644); err != nil {
			return err
		}
	}

	return nil
}
